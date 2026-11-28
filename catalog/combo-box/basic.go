package combobox

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/field"
)

func Basic(c *ui.Context) {
	root := ui.Column(c).Width(265).Gap(8)
	preset := ui.Local(root, "preset", func() string { return "Screen" })
	root.Children(func() {
		ui.Text(c, "Export preset").FontWeight(600)
		field.ComboBox(c, preset, field.ComboBoxProps{
			Label: "Export preset", Options: []string{"Screen", "Print", "Archive"},
		})
	})
}
