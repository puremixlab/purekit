package hotkey

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("hotkey/basic", Basic)
	registry.Add("hotkey/variants", Variants)
	registry.Add("hotkey/sizes", Sizes)
}
