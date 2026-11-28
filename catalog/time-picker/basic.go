package timepicker

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/field"
)

func Basic(c *ui.Context) {
	root := ui.Column(c).Gap(8)
	value := ui.Local(root, "time", func() field.TimeOfDay { return field.TimeOfDay{Hour: 9, Minute: 30} })
	root.Children(func() {
		ui.Text(c, "Reminder time").FontWeight(600)
		field.TimePicker(c, value, field.TimePickerProps{
			Label: "Reminder time", TwelveHour: true, Seconds: true,
		})
	})
}
