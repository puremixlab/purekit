package actionicon

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("action-icon/basic", Basic)
	registry.Add("action-icon/variants", Variants)
	registry.Add("action-icon/sizes", Sizes)
}
