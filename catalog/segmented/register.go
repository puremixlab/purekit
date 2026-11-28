package segmented

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("segmented/basic", Basic)
	registry.Add("segmented/variants", Variants)
}
