package columnview

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/display"
	"github.com/puremixlab/purekit/pkg/lucide"
)

var children = map[string][]string{
	"Workspace": {"Design", "Research", "Release"},
	"Design":    {"Brief.md", "Asset catalog.pdf", "Screens"},
	"Screens":   {"Overview.png", "Detail.png"},
	"Research":  {"Interviews.md", "References.csv"},
	"Release":   {"Notes.md"},
	"Archive":   {"2025"},
}

func Basic(c *ui.Context) {
	root := ui.Column(c).Fill()
	state := ui.Local(root, "columns", func() display.ColumnViewState[string] {
		return display.ColumnViewState[string]{Path: []string{"Workspace", "Design"}}
	})
	root.Children(func() {
		display.ColumnView(c, state, display.ColumnViewProps[string]{
			Label: "Folders", Roots: []string{"Workspace", "Archive"}, ColumnWidth: 158,
			Children: func(item string) []string { return children[item] },
		}, func(item string) {
			icon := lucide.File()
			if len(children[item]) > 0 {
				icon = lucide.Folder()
			}
			ui.Row(c).AlignItems(ui.Center).Gap(7).Children(func() {
				ui.Icon(c, icon).Size(14, 14)
				ui.Text(c, item).SingleLine()
			})
		})
	})
}
