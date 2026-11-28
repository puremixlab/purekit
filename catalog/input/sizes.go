package input

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/field"
	"github.com/puremixlab/purekit/pkg/theme"
)

var sizeValues = make([]string, 3)

func Sizes(c *ui.Context) {
	ui.Column(c).Gap(12).Width(320).Children(func() {
		field.Input(c, &sizeValues[0], field.InputProps{Placeholder: "Small", Size: theme.Small})
		field.Input(c, &sizeValues[1], field.InputProps{Placeholder: "Middle"})
		field.Input(c, &sizeValues[2], field.InputProps{Placeholder: "Large", Size: theme.Large})
	})
}
