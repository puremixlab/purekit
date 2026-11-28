package field

import (
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/internal/l10n"
	"github.com/puremixlab/purekit/pkg/lucide"
	"github.com/puremixlab/purekit/pkg/overlay"
	"github.com/puremixlab/purekit/pkg/theme"
)

type DatePickerProps struct {
	// Placeholder shows while no day is chosen: "Select date" when empty.
	Placeholder string
	// Format writes the chosen day in the trigger, as time.Format takes it: "2006-01-02" when
	// empty.
	Format string
	// AllowClear adds a button that empties the value.
	AllowClear bool
	// DisabledDate reports the days that cannot be chosen, as the Calendar's.
	DisabledDate func(day time.Time) bool
	// WeekStart is the first day of each week: Sunday when zero.
	WeekStart time.Weekday
	// NoToday hides the calendar's Today button.
	NoToday bool
	// Variant picks the trigger's surface, as Input's.
	Variant InputVariant
	// Size sets the trigger's height (24, 32 or 40), padding and font size, as Input's.
	Size theme.Size
	// Disabled keeps the picker from opening.
	Disabled bool
}

// DatePicker chooses a day into *value, the zero time while none is chosen, from a Calendar that
// opens below its trigger and closes once a day is chosen. It returns the trigger.
func DatePicker(c *ui.Context, value *time.Time, p DatePickerProps) ui.Element {
	if p.Format == "" {
		p.Format = "2006-01-02"
	}
	if p.Placeholder == "" {
		p.Placeholder = l10n.Text(c, "datePickerPlaceholder")
	}
	var label string
	if !value.IsZero() {
		label = value.Format(p.Format)
	}
	return dateControl(c, p, label, func() { *value = time.Time{} }, func() bool {
		_, picked := calendar(c, value, p.calendarProps(), nil)
		return picked
	})
}

// DateRange stores the endpoints of a date range. End is zero while an endpoint is pending.
type DateRange struct {
	Start, End time.Time
}

// DateRangePicker chooses a start and an end day using the same calendar as DatePicker.
// Choosing an earlier second endpoint arranges the two days in chronological order.
func DateRangePicker(c *ui.Context, value *DateRange, p DatePickerProps) ui.Element {
	if p.Format == "" {
		p.Format = "2006-01-02"
	}
	if p.Placeholder == "" {
		p.Placeholder = l10n.Text(c, "dateRangePlaceholder")
	}
	var label string
	if !value.Start.IsZero() {
		label = value.Start.Format(p.Format)
		if value.End.IsZero() {
			label += " – " + l10n.Text(c, "dateRangeEnd")
		} else {
			label += " – " + value.End.Format(p.Format)
		}
	}
	return dateControl(c, p, label, func() { *value = DateRange{} }, func() bool {
		day := value.Start
		if !value.End.IsZero() {
			day = value.End
		}
		_, picked := calendar(c, &day, p.calendarProps(), value)
		if !picked {
			return false
		}
		if value.Start.IsZero() || !value.End.IsZero() {
			*value = DateRange{Start: day}
			c.Invalidate()
			return false
		}
		value.End = day
		if day.Before(value.Start) {
			value.Start, value.End = day, value.Start
		}
		return true
	})
}

func (p DatePickerProps) calendarProps() CalendarProps {
	return CalendarProps{DisabledDate: p.DisabledDate, WeekStart: p.WeekStart, NoToday: p.NoToday}
}

func dateControl(c *ui.Context, p DatePickerProps, label string, clear func(), body func() bool) ui.Element {
	tk := theme.TokensOf(c)
	box := ui.Row(c).FillWidth().AlignItems(ui.Center)
	open := ui.Local(box, "open", func() bool { return false })
	var trigger ui.Element
	box.Children(func() {
		trigger = styleFieldRow(ui.ButtonBase(c), tk, p.Size).Grow(1).Gap(8).Cursor(ui.CursorPointer).FocusRing(false)
		if p.Disabled {
			trigger.Disabled(true).Cursor(ui.CursorNotAllowed)
			box.Opacity(0.5)
		}
		clearable := p.AllowClear && label != "" && !p.Disabled
		trigger.Children(func() {
			if label == "" {
				ui.Text(c, p.Placeholder).TextColor(tk.TextPlaceholder).Grow(1).SingleLine()
			} else {
				ui.Text(c, label).FontFeatures("tnum").Grow(1).SingleLine()
			}
			icon := ui.Icon(c, lucide.Calendar()).TextColor(tk.TextTertiary).Shrink(0)
			if clearable {
				icon.Opacity(0)
			}
		})
		styleField(trigger, tk, p.Variant, false, *open || trigger.Focused())
		if clearable {
			x := ui.ButtonBase(c).Label(l10n.Text(c, "clear")).Absolute().Right(10).Top(0).Bottom(0).Width(16).Radius(8).Center().
				Cursor(ui.CursorPointer).TextColor(tk.TextTertiary)
			if x.Hovered() {
				x.TextColor(tk.TextSecondary)
			}
			if x.Clicked() {
				clear()
			}
			x.Children(func() { ui.Icon(c, lucide.CircleX()).FontSize(14) })
		}
		if p.Disabled {
			return
		}
		picked := false
		panel := overlay.Popover(c, trigger, open, overlay.PopoverProps{Placement: overlay.Bottom}, func() {
			picked = body()
		})
		panel.Padding(0).Radius(16)
		if picked {
			*open = false
		}
	})
	return trigger
}
