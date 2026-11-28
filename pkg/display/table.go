package display

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/theme"
)

type TableProps struct {
	// Label is the accessible name of the table.
	Label string
	// Columns describe the headers, initial widths, alignment and sortable columns.
	Columns []ui.TableColumn
	// Rows is the number of rows. Only visible rows and a small margin are built.
	Rows int
}

// Table builds a data table with resizable columns and a scrolling viewport. Give each
// table its own ListState to keep selection, scrolling, sort order and column layout.
// When state.Sort or state.SortBy changes, apply that ordering to the data passed to cell.
func Table(c *ui.Context, state *ui.ListState, p TableProps, cell func(row, column int)) ui.Element {
	return styleDataView(c, ui.Table(c, state, p.Columns, p.Rows, cell), p.Label)
}

type OutlineProps[K comparable] struct {
	// Label is the accessible name of the outline.
	Label string
	// Roots are the top-level items. Each item key is unique throughout the hierarchy.
	Roots []K
	// Children returns the direct children of an item, or nil for a leaf. The hierarchy
	// must be acyclic. If Children is nil, all roots are leaves.
	Children func(item K) []K
}

// Outline renders a virtualized tree whose open branches and selected items are kept in state.
func Outline[K comparable](c *ui.Context, state *ui.OutlineState[K], p OutlineProps[K], row func(item K)) ui.Element {
	return styleDataView(c, ui.Outline(c, state, p.Roots, p.children, row), p.Label)
}

// OutlineTable adds columns to an outline. The first column shows hierarchy and disclosure
// controls. Keep each sibling group in the desired order when state.List.Sort changes.
func OutlineTable[K comparable](c *ui.Context, state *ui.OutlineState[K], p OutlineProps[K], columns []ui.TableColumn, cell func(item K, column int)) ui.Element {
	return styleDataView(c, ui.OutlineTable(c, state, columns, p.Roots, p.children, cell), p.Label)
}

func (p OutlineProps[K]) children(item K) []K {
	if p.Children == nil {
		return nil
	}
	return p.Children(item)
}

func styleDataView(c *ui.Context, view ui.Element, label string) ui.Element {
	tk := theme.TokensOf(c)
	return view.Label(label).MinWidth(0).MinHeight(0).AlignSelf(ui.Stretch).
		Background(tk.BgContainer).Border(1, tk.BorderSecondary).Radius(tk.BorderRadiusSM)
}
