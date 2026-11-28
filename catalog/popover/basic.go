package popover

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/action"
	"github.com/puremixlab/purekit/pkg/overlay"
	"github.com/puremixlab/purekit/pkg/theme"
)

var shareOpen bool

func Basic(c *ui.Context) {
	tk := theme.TokensOf(c)
	share := action.Button(c, "Share", action.ButtonProps{})
	overlay.Popover(c, share, &shareOpen, overlay.PopoverProps{Placement: overlay.Bottom}, func() {
		ui.Column(c).Gap(8).Width(240).Children(func() {
			ui.Text(c, "Share this page").FontWeight(600)
			ui.Text(c, "Anyone with the link can view it.").TextColor(tk.TextSecondary)
			if action.Button(c, "Copy link", action.ButtonProps{Variant: action.Primary, Size: theme.Small}).AlignSelf(ui.End).Clicked() {
				shareOpen = false
			}
		})
	})
}
