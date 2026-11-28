package selectdemo

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("select/basic", Basic)
	registry.Add("select/variants", Variants)
}
