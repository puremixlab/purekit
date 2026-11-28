package input

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("input/basic", Basic)
	registry.Add("input/variants", Variants)
	registry.Add("input/sizes", Sizes)
}
