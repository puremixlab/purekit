package tooltip

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/action"
	"github.com/puremixlab/purekit/pkg/overlay"
)

func Basic(c *ui.Context) {
	ui.Row(c).Gap(12).Children(func() {
		save := action.Button(c, "Save", action.ButtonProps{Variant: action.Primary})
		overlay.Tooltip(c, save, "Save the draft", overlay.TooltipProps{Hotkey: "mod+s"})
		share := action.Button(c, "Share", action.ButtonProps{})
		overlay.Tooltip(c, share, "Anyone with the link can view this page", overlay.TooltipProps{})
	})
}
