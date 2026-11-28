package display

import (
	"fmt"
	"slices"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/theme"
)

// ColumnViewState stores the selection path, from the root to the selected item.
type ColumnViewState[K comparable] struct {
	Path    []K
	columns []*browserColumn[K]
}

type browserColumn[K comparable] struct {
	list     ui.ListState
	items    []K
	selected int
}

type ColumnViewProps[K comparable] struct {
	// Label is the accessible name of the column browser.
	Label string
	// Roots and Children supply the hierarchy. Return nil for a leaf, and use unique keys.
	Roots    []K
	Children func(K) []K
	// ItemLabel provides an accessible name for each row.
	ItemLabel func(K) string
	// ColumnWidth is the column width, 180 DIPs when zero.
	ColumnWidth float32
	// OnActivate runs on Enter or a double click on an item.
	OnActivate func(K)
}

// ColumnView browses a hierarchy in adjacent columns. A branch shows its children in
// the next column. Left and Right move between columns and each column scrolls independently.
func ColumnView[K comparable](c *ui.Context, state *ColumnViewState[K], p ColumnViewProps[K], row func(K)) ui.Element {
	tk := theme.TokensOf(c)
	width := p.ColumnWidth
	if width <= 0 {
		width = 180
	}
	root := ui.ScrollHorizontal(c).Fill().Role(ui.RoleRegion).Label(p.Label).MinWidth(0).MinHeight(0)
	root.Children(func() {
		ui.Row(c).FillHeight().AlignItems(ui.Stretch).Children(func() {
			focusNext := false
			items := p.Roots
			depth := 0
			for len(items) > 0 {
				column := state.column(depth, p.ItemLabel)
				column.items = items
				column.selected = -1
				if depth < len(state.Path) {
					column.selected = slices.Index(items, state.Path[depth])
					if column.selected < 0 {
						state.Path = state.Path[:depth]
					}
				}
				if focusNext && column.selected < 0 {
					column.selected = 0
					state.Path = append(state.Path[:depth], items[0])
				}
				var list ui.Element
				ui.Box(c.Key(depth)).Width(width).Shrink(0).
					BorderWidth(0, 1, 0, 0).BorderColor(tk.BorderSecondary).Children(func() {
					list = ui.List(c, &column.list, len(items), func(index int) {
						item := column.items[index]
						ui.Row(c).AlignItems(ui.Center).Gap(5).Padding(5, 6).Children(func() {
							ui.Box(c).Grow(1).MinWidth(0).Children(func() { row(item) })
							if p.Children != nil && len(p.Children(item)) > 0 {
								ui.Text(c, "›").TextColor(tk.TextSecondary)
							}
						})
					}).Fill()
				})
				if focusNext {
					column.list.Focus(c)
				}
				if list.Changed() && column.selected >= 0 && column.selected < len(items) {
					item := items[column.selected]
					if depth >= len(state.Path) || state.Path[depth] != item {
						state.Path = append(state.Path[:depth], item)
						c.Invalidate()
					}
				}
				if depth > 0 && column.list.Shortcut(c, 0, ui.KeyLeft) {
					state.columns[depth-1].list.Focus(c)
				}
				depth++
				if column.selected < 0 || column.selected >= len(items) {
					break
				}
				item := items[column.selected]
				if p.OnActivate != nil && list.Submitted() {
					p.OnActivate(item)
				}
				if p.Children == nil {
					break
				}
				items = p.Children(item)
				focusNext = len(items) > 0 && column.list.Shortcut(c, 0, ui.KeyRight)
			}
			if depth < len(state.Path) {
				state.Path = state.Path[:depth]
			}
		})
	})
	return root
}

func (state *ColumnViewState[K]) column(depth int, label func(K) string) *browserColumn[K] {
	if depth == len(state.columns) {
		column := &browserColumn[K]{selected: -1}
		column.list.Selected = &column.selected
		column.list.Key = func(row int) any { return column.items[row] }
		state.columns = append(state.columns, column)
	}
	column := state.columns[depth]
	column.list.Label = func(row int) string {
		if label != nil {
			return label(column.items[row])
		}
		return fmt.Sprint(column.items[row])
	}
	return column
}
