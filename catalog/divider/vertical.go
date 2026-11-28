package divider

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/layout"
	"github.com/puremixlab/purekit/pkg/theme"
)

func Vertical(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Row(c).Gap(10).Children(func() {
		for i, label := range []string{"Edit", "View", "Help"} {
			if i > 0 {
				layout.Divider(c.Key(i), layout.DividerProps{Vertical: true})
			}
			ui.Text(c.Key(label), label).TextColor(tk.TextSecondary)
		}
	})
}
