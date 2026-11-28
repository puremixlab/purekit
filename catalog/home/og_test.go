package home

import (
	"testing"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/theme"
)

func ogTester() *ui.Tester {
	tt := ui.NewTester(func(c *ui.Context) {
		theme.Apply(c)
		OG(c)
	}, 1280, 640)
	tt.SetDark(true)
	tt.Frame()
	return tt
}

func TestOGLeftColumn(t *testing.T) {
	tt := ogTester()
	name, ok := tt.Find("PureKit")
	if !ok || name.X > 260 || name.Y > 160 {
		t.Errorf("name %+v, want at the top left after the marks", name)
	}
	first, ok1 := tt.Find(ogHeadline[0])
	second, ok2 := tt.Find(ogHeadline[1])
	if !ok1 || !ok2 || first.X != 64 || second.Y <= first.Y || first.H < 60 {
		t.Errorf("headline %+v / %+v, want two large lines at x 64", first, second)
	}
	desc, ok := tt.Find(ogDescription[0])
	if !ok || desc.Y < second.Y+second.H+40 || desc.X != 64 {
		t.Errorf("description %+v, want under the headline at x 64", desc)
	}
	if right := max(first.X+first.W, second.X+second.W); right > 680 {
		t.Errorf("headline reaches x %v, want it clear of the component preview", right)
	}
}

func TestOGPreviewSitsBesideTheHeadline(t *testing.T) {
	tt := ogTester()
	preview, ok := tt.Find("Components")
	if !ok || preview.X < 680 || preview.X+preview.W > 1280 {
		t.Errorf("component preview %+v, want it to the right of the headline", preview)
	}
}
