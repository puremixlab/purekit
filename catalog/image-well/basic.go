package imagewell

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/field"
)

func Basic(c *ui.Context) {
	root := ui.Column(c)
	value := ui.Local(root, "image", func() field.ImageWellValue { return field.ImageWellValue{} })
	root.Children(func() {
		field.ImageWell(c, value, field.ImageWellProps{})
	})
}
