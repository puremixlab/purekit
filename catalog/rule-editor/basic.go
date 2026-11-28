package ruleeditor

import (
	"strings"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/field"
)

func Basic(c *ui.Context) {
	root := ui.Column(c).FillWidth().Gap(11)
	rules := ui.Local(root, "rules", func() field.RuleSet {
		return field.RuleSet{Rules: []*field.Rule{{Field: "name", Operator: "contains", Value: ".md"}}}
	})
	root.Children(func() {
		field.RuleEditor(c, rules, field.RuleEditorProps{
			Fields: []field.RuleField{
				{Key: "name", Label: "Name", Operators: []field.RuleOperator{
					{Key: "contains", Label: "Contains"}, {Key: "is", Label: "Is"},
				}},
			},
		})
		for _, name := range []string{"Brief.md", "Inventory.csv", "Changes.md"} {
			if rules.Matches(func(rule field.Rule) bool {
				value := strings.ToLower(rule.Value)
				if rule.Operator == "is" {
					return strings.EqualFold(name, rule.Value)
				}
				return strings.Contains(strings.ToLower(name), value)
			}) {
				ui.Text(c.Key(name), name)
			}
		}
	})
}
