package outline

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/display"
	"github.com/puremixlab/purekit/pkg/lucide"
)

var children = map[string][]string{
	"Projects": {"Design", "Research"},
	"Design":   {"Brief.md", "Moodboard.png"},
	"Research": {"Notes.md", "Sources.csv"},
	"Archive":  {"2025", "2024"},
	"2025":     {"Release notes.md"},
}

func Basic(c *ui.Context) {
	root := ui.Column(c).Fill()
	state := ui.Local(root, "outline", func() ui.OutlineState[string] {
		var state ui.OutlineState[string]
		state.Open.Add("Projects")
		return state
	})
	selected := ui.Local(root, "selection", func() int { return -1 })
	state.List.Selected = selected
	root.Children(func() {
		display.Outline(c, state, display.OutlineProps[string]{
			Label: "Folders", Roots: []string{"Projects", "Archive"},
			Children: func(item string) []string { return children[item] },
		}, func(item string) {
			icon := lucide.File()
			if children[item] != nil {
				icon = lucide.Folder()
			}
			ui.Icon(c, icon).FontSize(15)
			ui.Text(c, item).SingleLine()
		}).Grow(1)
	})
}
