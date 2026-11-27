package display

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/lobehub/lobe-mygo/pkg/internal/testkit"
	"github.com/lobehub/lobe-mygo/pkg/theme"
)

func textTester(s string, p TextProps, width float32) *ui.Tester {
	return ui.NewTester(func(c *ui.Context) {
		theme.Apply(c)
		ui.Column(c).Padding(20).Children(func() {
			ui.Column(c).Width(width).Children(func() { Text(c, s, p) })
		})
	}, 400, 200)
}

func TestTextTypesColorTheInk(t *testing.T) {
	tk := theme.Light()
	for _, tc := range []struct {
		typ  TextType
		want ui.Color
	}{{TextDefault, tk.Text}, {TextSecondary, tk.TextDescription}, {TextDanger, tk.Error}, {TextSuccess, tk.Success}} {
		tt := textTester("MMMM", TextProps{Type: tc.typ, FontSize: 32, Strong: true}, 300)
		r, _ := tt.Find("MMMM")
		if got := testkit.Ink(tt, 400, r, tk.BgLayout); !testkit.Near(got, tc.want) {
			t.Errorf("type %d ink %+v, want %+v", tc.typ, got, tc.want)
		}
	}
	tt := textTester("MMMM", TextProps{Type: TextDanger, Disabled: true, FontSize: 32}, 300)
	r, _ := tt.Find("MMMM")
	if got := testkit.Ink(tt, 400, r, tk.BgLayout); !testkit.Near(got, tk.TextDisabled) {
		t.Errorf("disabled ink %+v, want TextDisabled", got)
	}
}

func TestTextHeadingsScaleTheFont(t *testing.T) {
	const title = "Title"
	const width = 300
	for i, scale := range []float32{2.5, 2, 1.5, 1.25, 1} {
		heading := i + 1
		tt := textTester(title, TextProps{Heading: heading}, width)
		got, ok := tt.Find(title)
		if !ok {
			t.Fatalf("heading %d did not render", heading)
		}

		// Shaping accounts for the platform font's ascent and descent.
		font := ui.Font{Size: theme.Light().FontSize * scale, Weight: 700, LineHeight: 1.25}
		_, want := ui.ShapeText(title, font, width).Size()
		if got.H < want-0.1 || got.H > want+0.1 {
			t.Errorf("heading %d line box %.2f, want %.2f", heading, got.H, want)
		}
	}
}

func TestTextEllipsisRows(t *testing.T) {
	long := "Talk to your agents and keep every conversation in one place, searchable and shared."
	one := textTester(long, TextProps{Rows: 1}, 120)
	two := textTester(long, TextProps{Rows: 2}, 120)
	free := textTester(long, TextProps{}, 120)
	r1, _ := one.Find(long)
	r2, _ := two.Find(long)
	rf, _ := free.Find(long)
	if !(r1.H < r2.H && r2.H < rf.H) || r2.H > 2*r1.H+1 {
		t.Errorf("heights one %v two %v free %v: want one line, two lines, more", r1.H, r2.H, rf.H)
	}
}

func TestTextMarkHighlights(t *testing.T) {
	tk := theme.Light()
	tt := textTester("Note", TextProps{Mark: true}, 300)
	r, _ := tt.Find("Note")
	if got := testkit.PixelAt(tt, 400, r.X+r.W/2, r.Y+1); !testkit.Near(got, tk.Yellow) {
		t.Errorf("mark background %+v, want Yellow", got)
	}
}

func TestTextShinySweeps(t *testing.T) {
	tk := theme.Light()
	tt := textTester("MMMMMMMMMMMM", TextProps{Shiny: true, FontSize: 32, Strong: true}, 380)
	r, _ := tt.Find("MMMMMMMMMMMM")
	left := ui.Rect{X: r.X, Y: r.Y, W: r.W / 4, H: r.H}
	right := ui.Rect{X: r.X + r.W*3/4, Y: r.Y, W: r.W / 4, H: r.H}
	ink := func(rect ui.Rect) int {
		c := testkit.Ink(tt, 400, rect, tk.BgLayout)
		return int(tk.BgLayout.R) - int(c.R)
	}
	tt.Advance(225 * time.Millisecond)
	tt.Frame()
	l1, r1 := ink(left), ink(right)
	tt.Advance(300 * time.Millisecond)
	tt.Frame()
	l2, r2 := ink(left), ink(right)
	if !(l1 > r1 && r2 > l2) {
		t.Errorf("sweep: at 225 ms left %d right %d, at 525 ms left %d right %d; want the bright band moving left to right, as CSS's background-position 100%% to -100%% does", l1, r1, l2, r2)
	}
}
