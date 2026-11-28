package button

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("button/basic", Basic)
	registry.Add("button/variants", Variants)
	registry.Add("button/sizes", Sizes)
}
