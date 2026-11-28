package list

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("list/basic", Basic)
	registry.Add("list/variants", Variants)
}
