package copybutton

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/action"
	"github.com/puremixlab/purekit/pkg/theme"
)

func Basic(c *ui.Context) {
	ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
		ui.Text(c, "go get github.com/puremixlab/purekit").Font("monospace").FontSize(13)
		action.CopyButton(c, "go get github.com/puremixlab/purekit", action.CopyButtonProps{})
		action.CopyButton(c.Key("filled"), "go get github.com/puremixlab/purekit", action.CopyButtonProps{Variant: action.ActionIconFilled, Size: theme.Small})
	})
}
