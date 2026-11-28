package field

import (
	"math"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/internal/l10n"
	"github.com/puremixlab/purekit/pkg/theme"
)

// TimeOfDay is a clock time without a date or time zone.
type TimeOfDay struct {
	Hour, Minute, Second int
}

type TimePickerProps struct {
	// Label is the accessible name of the group of fields.
	Label string
	// TwelveHour displays a 1–12 hour field with an AM or PM control.
	TwelveHour bool
	// Seconds shows a field for seconds in addition to hours and minutes.
	Seconds bool
	// Disabled prevents changes to the time.
	Disabled bool
}

// TimePicker edits hours, minutes and optionally seconds. The stored hour is always
// on a 24-hour clock, including when the display uses AM and PM.
func TimePicker(c *ui.Context, value *TimeOfDay, p TimePickerProps) ui.Element {
	label := p.Label
	if label == "" {
		label = l10n.Text(c, "timePicker")
	}
	root := ui.Row(c).Gap(5).AlignItems(ui.Center).Wrap().Role(ui.RoleGroup).Label(label).Disabled(p.Disabled)
	root.Children(func() {
		hour := min(max(value.Hour, 0), 23)
		period := "am"
		if hour >= 12 {
			period = "pm"
		}
		if p.TwelveHour {
			shown := (hour+11)%12 + 1
			hour = timeNumber(c.Key("hour"), shown, 1, 12, l10n.Text(c, "timeHour")) % 12
		} else {
			hour = timeNumber(c.Key("hour"), hour, 0, 23, l10n.Text(c, "timeHour"))
		}
		ui.Text(c, ":")
		value.Minute = timeNumber(c.Key("minute"), value.Minute, 0, 59, l10n.Text(c, "timeMinute"))
		if p.Seconds {
			ui.Text(c, ":")
			value.Second = timeNumber(c.Key("second"), value.Second, 0, 59, l10n.Text(c, "timeSecond"))
		}
		if p.TwelveHour {
			Select(c, &period, SelectProps{
				Size: theme.Small, Disabled: p.Disabled,
				Options: []SelectOption{
					{Value: "am", Label: l10n.Text(c, "timeAM")},
					{Value: "pm", Label: l10n.Text(c, "timePM")},
				},
			}).Label(l10n.Text(c, "timePeriod")).Width(74)
			if period == "pm" {
				hour += 12
			}
		}
		value.Hour = hour
	})
	return root
}

func timeNumber(c *ui.Context, initial, lo, hi int, label string) int {
	initial = min(max(initial, lo), hi)
	value := float64(initial)
	ui.Box(c).Width(64).Children(func() {
		InputNumber(c, &value, InputNumberProps{
			Min: float64(lo), Max: float64(hi), Step: 1, NoControls: true, Size: theme.Small,
		}).Label(label)
	})
	if math.IsNaN(value) {
		return initial
	}
	return int(value)
}
