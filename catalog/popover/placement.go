package popover

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/action"
	"github.com/puremixlab/purekit/pkg/overlay"
)

var placementOpen = map[overlay.Placement]*bool{
	overlay.Top: new(bool), overlay.Bottom: new(bool), overlay.Left: new(bool), overlay.Right: new(bool),
}

func Placement(c *ui.Context) {
	sides := []struct {
		label     string
		placement overlay.Placement
	}{{"Top", overlay.Top}, {"Bottom", overlay.Bottom}, {"Left", overlay.Left}, {"Right", overlay.Right}}
	ui.Row(c).Gap(12).Children(func() {
		for _, s := range sides {
			b := action.Button(c.Key(s.label), s.label, action.ButtonProps{})
			overlay.Popover(c, b, placementOpen[s.placement], overlay.PopoverProps{Placement: s.placement}, func() {
				ui.Text(c, "Opens on the "+s.label)
			})
		}
	})
}
