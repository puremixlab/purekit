package home

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/theme"
)

func TestBannerIsARoundedCardWithTheHeadline(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		theme.Apply(c)
		ui.Box(c).Fill().Background(ui.Hex("#FFFFFF")).Children(func() { Banner(c) })
	}, 1474, 602)
	tt.SetDark(true)
	tt.Frame()
	preview, ok := tt.Find("Components")
	if !ok || preview.X < 680 || preview.X+preview.W > 1474 {
		t.Errorf("component preview %+v, want it to the right of the headline", preview)
	}
	head, ok := tt.Find(ogHeadline[0])
	if !ok || head.X != 64 || head.H < 60 {
		t.Errorf("headline %+v, want a large line at x 64", head)
	}
	img := tt.Image()
	if c := img.RGBAAt(0, 0); c.R < 240 || c.G < 240 || c.B < 240 {
		t.Errorf("corner pixel %+v, want the page showing through the rounded corner", c)
	}
	if c := img.RGBAAt(30, 300); c.R > 40 || c.G > 40 || c.B > 40 {
		t.Errorf("left edge pixel %+v, want the dark card", c)
	}
}
