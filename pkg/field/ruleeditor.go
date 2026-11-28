package field

import (
	"slices"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/action"
	"github.com/puremixlab/purekit/pkg/internal/l10n"
	"github.com/puremixlab/purekit/pkg/lucide"
	"github.com/puremixlab/purekit/pkg/theme"
)

type RuleOperator struct {
	// Key is the value stored in a Rule. Label is the text shown in its selector.
	Key, Label string
	// NoValue hides the right-hand operand for operations such as "is empty".
	NoValue bool
}

type RuleField struct {
	// Key identifies the attribute to test. Label is its visible name.
	Key, Label string
	// Operators are the comparisons available for this attribute.
	Operators []RuleOperator
}

type Rule struct {
	Field, Operator, Value string
}

// RuleSet is an editable group of comparisons. Keep each rule's pointer while it is in the set.
type RuleSet struct {
	// Any matches when at least one rule matches. Otherwise all rules must match.
	Any   bool
	Rules []*Rule
}

// Matches combines the results of the app's predicate. An empty set matches everything.
func (set RuleSet) Matches(predicate func(Rule) bool) bool {
	hasRule := false
	for _, rule := range set.Rules {
		if rule == nil {
			continue
		}
		hasRule = true
		matches := predicate(*rule)
		if set.Any && matches {
			return true
		}
		if !set.Any && !matches {
			return false
		}
	}
	return !set.Any || !hasRule
}

type RuleEditorProps struct {
	// Label is the accessible name of the editor.
	Label string
	// Fields maps the app's attributes to their allowed comparisons.
	Fields []RuleField
	// Disabled makes the group read-only.
	Disabled bool
}

// RuleEditor builds rows of attribute, operator and value fields, with add and remove actions
// and a choice between matching all or any of the rules.
func RuleEditor(c *ui.Context, set *RuleSet, p RuleEditorProps) ui.Element {
	label := p.Label
	if label == "" {
		label = l10n.Text(c, "ruleEditor")
	}
	root := ui.Column(c).FillWidth().Gap(8).Role(ui.RoleGroup).Label(label).Disabled(p.Disabled)
	fields := make([]SelectOption, 0, len(p.Fields))
	for _, field := range p.Fields {
		fields = append(fields, SelectOption{Value: field.Key, Label: field.Label})
	}
	root.Children(func() {
		mode := "all"
		if set.Any {
			mode = "any"
		}
		ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
			Select(c, &mode, SelectProps{
				Disabled: p.Disabled, Size: theme.Small,
				Options: []SelectOption{
					{Value: "all", Label: l10n.Text(c, "ruleAll")},
					{Value: "any", Label: l10n.Text(c, "ruleAny")},
				},
			}).Label(l10n.Text(c, "ruleMatchMode"))
			set.Any = mode == "any"
			add := action.Button(c, l10n.Text(c, "ruleAdd"), action.ButtonProps{Size: theme.Small}).
				Disabled(p.Disabled || len(p.Fields) == 0)
			if add.Clicked() && !p.Disabled && len(p.Fields) > 0 {
				field := p.Fields[0]
				rule := &Rule{Field: field.Key}
				if len(field.Operators) > 0 {
					rule.Operator = field.Operators[0].Key
				}
				set.Rules = append(set.Rules, rule)
			}
		})
		var removed *Rule
		for _, rule := range set.Rules {
			if rule != nil && ruleEditorRow(c.Key(rule), rule, p, fields) {
				removed = rule
			}
		}
		if removed != nil {
			set.Rules = slices.DeleteFunc(set.Rules, func(rule *Rule) bool { return rule == removed })
			c.Invalidate()
		}
	})
	return root
}

func ruleEditorRow(c *ui.Context, rule *Rule, p RuleEditorProps, fields []SelectOption) bool {
	remove := false
	ui.Row(c).Gap(6).AlignItems(ui.Center).Wrap().Children(func() {
		oldField := rule.Field
		Select(c, &rule.Field, SelectProps{
			Options: fields, Placeholder: l10n.Text(c, "ruleField"), Size: theme.Small, Disabled: p.Disabled,
		}).Label(l10n.Text(c, "ruleField")).Width(122)
		var operators []RuleOperator
		for _, field := range p.Fields {
			if field.Key == rule.Field {
				operators = field.Operators
				break
			}
		}
		if rule.Field != oldField {
			rule.Operator, rule.Value = "", ""
			if len(operators) > 0 {
				rule.Operator = operators[0].Key
			}
		}
		options := make([]SelectOption, 0, len(operators))
		for _, op := range operators {
			options = append(options, SelectOption{Value: op.Key, Label: op.Label})
		}
		Select(c, &rule.Operator, SelectProps{
			Options: options, Placeholder: l10n.Text(c, "ruleOperator"), Size: theme.Small, Disabled: p.Disabled,
		}).Label(l10n.Text(c, "ruleOperator")).Width(126)
		valueNeeded := !slices.ContainsFunc(operators, func(op RuleOperator) bool {
			return op.Key == rule.Operator && op.NoValue
		})
		if valueNeeded {
			ui.Box(c).MinWidth(92).Grow(1).Children(func() {
				Input(c, &rule.Value, InputProps{Placeholder: l10n.Text(c, "ruleValue"), Size: theme.Small}).
					Label(l10n.Text(c, "ruleValue"))
			})
		}
		remove = action.ActionIcon(c, lucide.Minus(), action.ActionIconProps{
			Title: l10n.Text(c, "ruleRemove"), Size: theme.Small,
		}).Disabled(p.Disabled).Clicked() && !p.Disabled
	})
	return remove
}
