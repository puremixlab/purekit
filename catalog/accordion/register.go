package accordion

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("accordion/basic", Basic)
	registry.Add("accordion/variants", Variants)
}
