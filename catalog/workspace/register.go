package workspace

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/catalog/registry"
)

func init() {
	registry.Add("workspace/basic", func(c *ui.Context) {
		root := ui.Box(c).Fill()
		state := ui.Local(root, "workspace", New)
		root.Children(func() { (*state).View(c) })
	})
}
