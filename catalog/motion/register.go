package motion

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("motion/springs", Springs)
	registry.Add("motion/presets", Presets)
}
