package field

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestComboBoxShowsEveryOptionAndAcceptsCustomText(t *testing.T) {
	value := "Screen"
	tester := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Width(240).Children(func() {
			ComboBox(c, &value, ComboBoxProps{Label: "Preset", Options: []string{"Screen", "Print"}})
		})
	}, 400, 360)
	clickRole(t, tester, ui.RoleButton, "Show options")
	tester.Advance(time.Second)
	if !tester.HasText("Print") {
		t.Fatal("showing the options filtered them by the current value")
	}
	if err := tester.Click("Print"); err != nil {
		t.Fatal(err)
	}
	if value != "Print" {
		t.Fatalf("choice = %q, want Print", value)
	}
	if err := tester.Click("Preset"); err != nil {
		t.Fatal(err)
	}
	tester.Key(ui.Cmd, ui.KeyA)
	tester.Type("Personal")
	if value != "Personal" {
		t.Fatalf("custom value = %q, want Personal", value)
	}
}

func TestDateRangePickerAcceptsEndpointsInEitherOrder(t *testing.T) {
	var value DateRange
	tester := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Width(280).Children(func() {
			DateRangePicker(c, &value, DatePickerProps{AllowClear: true}).Label("Range")
		})
	}, 420, 480)
	clickRole(t, tester, ui.RoleButton, "Range")
	tester.Advance(time.Second)
	now := tester.Now()
	day := func(d int) time.Time { return time.Date(now.Year(), now.Month(), d, 0, 0, 0, 0, time.Local) }
	clickRole(t, tester, ui.RoleButton, day(18).Format("2006-01-02"))
	if !value.Start.Equal(day(18)) || !value.End.IsZero() {
		t.Fatalf("after the first choice got %+v", value)
	}
	clickRole(t, tester, ui.RoleButton, day(2).Format("2006-01-02"))
	tester.Advance(time.Second)
	if !value.Start.Equal(day(2)) || !value.End.Equal(day(18)) {
		t.Fatalf("range did not order its endpoints: %+v", value)
	}
	clickRole(t, tester, ui.RoleButton, "Clear")
	if !value.Start.IsZero() || !value.End.IsZero() {
		t.Fatalf("clear left a selected endpoint: %+v", value)
	}
}

func TestTimePickerKeepsClockValueInBounds(t *testing.T) {
	value := TimeOfDay{Hour: 14, Minute: 58, Second: 21}
	tester := ui.NewTester(func(c *ui.Context) {
		TimePicker(c, &value, TimePickerProps{TwelveHour: true})
	}, 420, 240)
	if err := tester.Click("Minutes"); err != nil {
		t.Fatal(err)
	}
	tester.Key(0, ui.KeyUp)
	tester.Key(0, ui.KeyUp)
	tester.Frame()
	if value != (TimeOfDay{Hour: 14, Minute: 59, Second: 21}) {
		t.Fatalf("time after incrementing minutes: %+v", value)
	}
}
