package field

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/richtext"
)

func TestRuleEditorAddsEditsAndRemovesRules(t *testing.T) {
	var set RuleSet
	tester := ui.NewTester(func(c *ui.Context) {
		RuleEditor(c, &set, RuleEditorProps{
			Fields: []RuleField{{Key: "name", Label: "Name", Operators: []RuleOperator{{Key: "is", Label: "Is"}}}},
		})
	}, 520, 255)
	clickRole(t, tester, ui.RoleButton, "Add rule")
	if len(set.Rules) != 1 || set.Rules[0].Field != "name" || set.Rules[0].Operator != "is" {
		t.Fatalf("unexpected new rule: %v", set.Rules)
	}
	if err := tester.Click("Value"); err != nil {
		t.Fatal(err)
	}
	tester.Type("Desktop")
	if !set.Matches(func(rule Rule) bool { return rule.Value == "Desktop" }) {
		t.Fatal("editing the operand did not change the active rules")
	}
	clickRole(t, tester, ui.RoleButton, "Remove rule")
	if len(set.Rules) != 0 || !set.Matches(func(Rule) bool { return false }) {
		t.Fatal("removing the rule did not reset matching")
	}
}

func TestRichTextToolbarFormatsTheSelection(t *testing.T) {
	doc := richtext.NewDocument(richtext.NewText("Selected text"))
	doc.SetSelection(0, 8)
	edits := 0
	tester := ui.NewTester(func(c *ui.Context) {
		RichTextEditor(c, doc, RichTextEditorProps{OnChange: func(richtext.Text) { edits++ }})
	}, 620, 260)
	if err := tester.Click("Bold"); err != nil {
		t.Fatal(err)
	}
	tester.Frame()
	if doc.Formatting().Bold != richtext.FormatOn || edits == 0 {
		t.Fatal("toolbar formatting did not update the selected text")
	}
}
