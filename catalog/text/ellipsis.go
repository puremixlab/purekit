package text

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/display"
)

const paragraph = "PureKit is an open-source component library for native Go apps. It includes controls for form input, dialogs, and navigation, with light and dark themes."

func Ellipsis(c *ui.Context) {
	ui.Column(c).Gap(16).Width(360).Children(func() {
		display.Text(c.Key(1), paragraph, display.TextProps{Rows: 1})
		display.Text(c.Key(2), paragraph, display.TextProps{Rows: 2, Type: display.TextSecondary})
	})
}
