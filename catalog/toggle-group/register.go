package togglegroup

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("toggle-group/basic", Basic)
	registry.Add("toggle-group/variants", Variants)
}
