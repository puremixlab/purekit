package actionicon

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/action"
	"github.com/puremixlab/purekit/pkg/lucide"
	"github.com/puremixlab/purekit/pkg/theme"
)

func Sizes(c *ui.Context) {
	ui.Row(c).Gap(12).AlignItems(ui.Center).Children(func() {
		for _, size := range []theme.Size{theme.Small, theme.Middle, theme.Large} {
			action.ActionIcon(c, lucide.Bell(), action.ActionIconProps{Title: "Notifications", Size: size, Variant: action.ActionIconFilled})
		}
	})
}
