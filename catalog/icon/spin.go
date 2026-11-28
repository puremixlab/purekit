package icon

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/display"
	"github.com/puremixlab/purekit/pkg/lucide"
)

func Spin(c *ui.Context) {
	ui.Row(c).Gap(16).Children(func() {
		display.Icon(c, lucide.LoaderCircle(), display.IconProps{Size: display.IconLarge, Spin: true})
		display.Icon(c, lucide.RefreshCw(), display.IconProps{Size: display.IconLarge, Spin: true})
	})
}
