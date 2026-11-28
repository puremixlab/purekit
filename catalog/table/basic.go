package table

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/display"
)

type document struct {
	name, kind string
	size       int
}

var documents = []document{
	{"Project brief.md", "Markdown", 18}, {"Release notes.md", "Markdown", 27},
	{"Budget.csv", "Spreadsheet", 43}, {"Overview.pdf", "PDF", 128},
	{"Timeline.csv", "Spreadsheet", 56}, {"Meeting notes.md", "Markdown", 12},
	{"Press kit.pdf", "PDF", 272}, {"Inventory.csv", "Spreadsheet", 38},
	{"Shipping guide.pdf", "PDF", 174}, {"Ideas.md", "Markdown", 9},
}

var columns = []ui.TableColumn{
	{ID: "name", Title: "Name", MinWidth: 144, Sortable: true},
	{ID: "kind", Title: "Kind", Width: 120, Sortable: true},
	{ID: "size", Title: "Size", Width: 96, Align: ui.End, Sortable: true},
}

type tableState struct {
	list      ui.ListState
	selected  int
	selection ui.Selection[string]
	sort      ui.SortOrder
}

func Basic(c *ui.Context) {
	root := ui.Column(c).Fill()
	state := ui.Local(root, "table", func() tableState { return tableState{selected: -1} })
	rows := slices.Clone(documents)
	order := state.sort
	slices.SortStableFunc(rows, func(a, b document) int {
		var comparison int
		switch order.Column {
		case "size":
			comparison = cmp.Compare(a.size, b.size)
		case "kind":
			comparison = cmp.Compare(a.kind, b.kind)
		default:
			comparison = cmp.Compare(a.name, b.name)
		}
		if order.Descending {
			return -comparison
		}
		return comparison
	})
	state.list.Selected = &state.selected
	state.list.Selection = &state.selection
	state.list.Sort = &state.sort
	state.list.Key = func(row int) any { return rows[row].name }
	state.list.Label = func(row int) string { return rows[row].name }
	root.Children(func() {
		display.Table(c, &state.list, display.TableProps{Label: "Documents", Columns: columns, Rows: len(rows)}, func(row, col int) {
			doc := rows[row]
			switch col {
			case 0:
				ui.Text(c, doc.name).SingleLine()
			case 1:
				ui.Text(c, doc.kind).SingleLine()
			case 2:
				ui.Text(c, fmt.Sprintf("%d KB", doc.size)).FontFeatures("tnum")
			}
		}).Grow(1)
	})
	if state.sort != order {
		c.Invalidate()
	}
}
