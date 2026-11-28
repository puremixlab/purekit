package progress

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("progress/basic", Basic)
	registry.Add("progress/variants", Variants)
}
