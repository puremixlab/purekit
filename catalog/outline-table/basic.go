package outlinetable

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/display"
)

var rows = map[string][]string{
	"Workspace": {"Specs", "Design"},
	"Specs":     {"Brief.md", "Scope.md"},
	"Design":    {"Screens.pdf"},
}

var columns = []ui.TableColumn{
	{ID: "name", Title: "Name", MinWidth: 155},
	{ID: "kind", Title: "Kind", Width: 120},
}

func Basic(c *ui.Context) {
	root := ui.Column(c).Fill()
	state := ui.Local(root, "outline", func() ui.OutlineState[string] {
		var state ui.OutlineState[string]
		state.Open.Add("Workspace")
		state.Open.Add("Specs")
		return state
	})
	state.List.Selected = ui.Local(root, "selected", func() int { return -1 })
	root.Children(func() {
		display.OutlineTable(c, state, display.OutlineProps[string]{
			Label: "Workspace files", Roots: []string{"Workspace"},
			Children: func(item string) []string { return rows[item] },
		}, columns, func(item string, column int) {
			if column == 0 {
				ui.Text(c, item).SingleLine()
			} else if rows[item] != nil {
				ui.Text(c, "Folder")
			} else {
				ui.Text(c, "Document")
			}
		}).Grow(1)
	})
}
