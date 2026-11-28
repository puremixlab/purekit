package skeleton

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("skeleton/basic", Basic)
	registry.Add("skeleton/animation", Animation)
}
