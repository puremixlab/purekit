package pathcontrol

import (
	"path/filepath"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/nav"
)

func Basic(c *ui.Context) {
	root := ui.Column(c).Gap(10)
	path := ui.Local(root, "path", func() string {
		return filepath.Join("Workspace", "Design", "Library", "Desktop", "Icons", "Navigation")
	})
	root.Children(func() {
		nav.PathControl(c, nav.PathControlProps{
			Path: *path, MaxVisible: 3, OnNavigate: func(parent string) { *path = parent },
		})
		ui.Text(c, *path).Wrap()
	})
}
