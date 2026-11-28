package tag

import (
	"fmt"
	"github.com/puremixlab/purekit/pkg/display"
	"github.com/puremixlab/purekit/pkg/theme"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/lucide"
)

var closed int

func Basic(c *ui.Context) {
	ui.Row(c).Gap(8).Children(func() {
		display.Tag(c, "Default", display.TagProps{})
		display.Tag(c, "With icon", display.TagProps{Icon: lucide.Sparkles(), Color: display.TagPurple})
		display.Tag(c, "Closable", display.TagProps{Closable: true, OnClose: func() { closed++ }})
	})
	ui.Text(c, fmt.Sprintf("Closed %d tags", closed)).TextColor(theme.TokensOf(c).TextSecondary)
}
