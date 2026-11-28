package button

import (
	"fmt"
	"github.com/puremixlab/purekit/pkg/action"
	"github.com/puremixlab/purekit/pkg/theme"

	"github.com/egoist/mygo/ui"
)

var saved int

func Basic(c *ui.Context) {
	ui.Row(c).Gap(8).Children(func() {
		if action.Button(c, "Reset", action.ButtonProps{}).Clicked() {
			saved = 0
		}
		if action.Button(c, "Save", action.ButtonProps{Variant: action.Primary}).Clicked() {
			saved++
		}
	})
	ui.Text(c, fmt.Sprintf("Saved %d times", saved)).TextColor(theme.TokensOf(c).TextSecondary)
}
