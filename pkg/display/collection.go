package display

import (
	"fmt"
	"slices"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/theme"
)

// CollectionState remembers the selection and scroll position of a collection.
type CollectionState[K comparable] struct {
	Selection  ui.Selection[K]
	list       ui.ListState
	current    K
	pivot      K
	hasCurrent bool
	hasPivot   bool
}

type CollectionViewProps[K comparable] struct {
	// Label is the accessible name of the collection.
	Label string
	// Items are stable, distinct keys for the tiles, in display order.
	Items []K
	// ItemLabel is the accessible name of a tile. By default it is the item key as text.
	ItemLabel func(K) string
	// Columns is the number of tiles across the view, 3 when zero.
	Columns int
	// ItemHeight is the height of a tile in DIPs, 84 when zero.
	ItemHeight float32
	// Multiple enables Command or Ctrl toggling and Shift range selection.
	Multiple bool
	// OnActivate runs on Enter or a double click.
	OnActivate func(K)
}

// CollectionView shows selectable tiles and builds only the rows near its viewport.
// Arrow keys and Home and End move the selection while the collection has focus.
func CollectionView[K comparable](c *ui.Context, state *CollectionState[K], p CollectionViewProps[K], tile func(K)) ui.Element {
	tk := theme.TokensOf(c)
	cols := p.Columns
	if cols <= 0 {
		cols = 3
	}
	height := p.ItemHeight
	if height <= 0 {
		height = 84
	}
	root := ui.Box(c).Role(ui.RoleListBox).Label(p.Label).Focusable().Fill().
		MinWidth(0).MinHeight(0).Background(tk.BgContainer)
	index := -1
	if state.hasCurrent {
		index = slices.Index(p.Items, state.current)
	}
	if len(p.Items) > 0 && index < 0 {
		index = 0
	}
	if index >= 0 {
		for _, movement := range []struct {
			key   ui.Key
			index int
		}{
			{ui.KeyLeft, index - 1}, {ui.KeyRight, index + 1},
			{ui.KeyUp, index - cols}, {ui.KeyDown, index + cols},
			{ui.KeyHome, 0}, {ui.KeyEnd, len(p.Items) - 1},
		} {
			mods := ui.Modifiers(0)
			if p.Multiple && root.Shortcut(ui.Shift, movement.key) {
				mods = ui.Shift
			} else if !root.Shortcut(0, movement.key) {
				continue
			}
			index = min(max(movement.index, 0), len(p.Items)-1)
			selectTile(state, p, index, mods)
			state.list.ScrollIntoView(index / cols)
			break
		}
		if root.Shortcut(0, ui.KeySpace) {
			selectTile(state, p, index, 0)
		}
		if p.OnActivate != nil && root.Shortcut(0, ui.KeyEnter) {
			p.OnActivate(p.Items[index])
		}
	}
	root.Children(func() {
		rows := (len(p.Items) + cols - 1) / cols
		ui.List(c, &state.list, rows, func(row int) {
			ui.Row(c).FillWidth().Gap(8).Padding(4).Children(func() {
				for col := range cols {
					i := row*cols + col
					if i >= len(p.Items) {
						ui.Box(c).Grow(1).Basis(0)
						continue
					}
					item := p.Items[i]
					label := fmt.Sprint(item)
					if p.ItemLabel != nil {
						label = p.ItemLabel(item)
					}
					b := ui.ButtonBase(c.Key(item)).Role(ui.RoleOption).Label(label).
						TabStop(false).KeepFocus().Grow(1).Basis(0).MinWidth(0).Height(height).
						Radius(tk.BorderRadiusSM).Center().Clip().Cursor(ui.CursorPointer).
						PositionInSet(i+1, len(p.Items)).Background(tk.FillQuaternary)
					if b.Clicked() {
						root.Focus()
						selectTile(state, p, i, b.ClickModifiers())
					}
					if p.OnActivate != nil && b.DoubleClicked() {
						p.OnActivate(item)
					}
					selected := state.Selection.Has(item)
					b.AccessSelected(selected)
					if selected {
						b.Background(tk.PrimaryBg).Border(1, tk.Primary)
					} else if b.Hovered() {
						b.Background(tk.FillTertiary)
					}
					if state.hasCurrent && state.current == item || !state.hasCurrent && i == index {
						root.ActiveDescendant(b)
					}
					b.Children(func() { tile(item) })
				}
			})
		}).Fill().Role(ui.RoleNone).TabStop(false)
	})
	return root
}

func selectTile[K comparable](state *CollectionState[K], p CollectionViewProps[K], index int, mods ui.Modifiers) {
	key := p.Items[index]
	anchor := -1
	if p.Multiple && mods&ui.Shift != 0 && state.hasPivot {
		anchor = slices.Index(p.Items, state.pivot)
	}
	switch {
	case anchor >= 0:
		state.Selection.Clear()
		for _, item := range p.Items[min(anchor, index) : max(anchor, index)+1] {
			state.Selection.Add(item)
		}
	case p.Multiple && mods&ui.Cmd != 0:
		if state.Selection.Has(key) {
			state.Selection.Remove(key)
		} else {
			state.Selection.Add(key)
		}
		state.pivot, state.hasPivot = key, true
	default:
		state.Selection.Clear()
		state.Selection.Add(key)
		state.pivot, state.hasPivot = key, true
	}
	state.current, state.hasCurrent = key, true
}
