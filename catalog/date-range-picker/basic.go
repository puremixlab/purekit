package daterangepicker

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/field"
)

func Basic(c *ui.Context) {
	root := ui.Column(c).Width(285).Gap(8)
	value := ui.Local(root, "range", func() field.DateRange {
		start := c.Now().AddDate(0, 0, -4)
		return field.DateRange{Start: start, End: c.Now()}
	})
	root.Children(func() {
		ui.Text(c, "Report dates").FontWeight(600)
		field.DateRangePicker(c, value, field.DatePickerProps{AllowClear: true})
	})
}
