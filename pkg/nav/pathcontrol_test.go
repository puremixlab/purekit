package nav

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func TestPathControlShowsParentsForRelativePath(t *testing.T) {
	path := filepath.Join("parents", "project", "assets", "icons")
	var selected string
	tester := ui.NewTester(func(c *ui.Context) {
		PathControl(c, PathControlProps{
			Path: path, MaxVisible: 2,
			OnNavigate: func(path string) { selected = path },
		})
	}, 440, 280)

	if err := tester.Click("Show parent folders"); err != nil {
		t.Fatal(err)
	}
	tester.Advance(time.Second)
	if err := tester.Click("project"); err != nil {
		t.Fatal(err)
	}
	tester.Frame()
	if want := filepath.Join("parents", "project"); selected != want {
		t.Errorf("selected path = %q, want %q", selected, want)
	}
}
