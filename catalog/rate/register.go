package rate

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("rate/basic", Basic)
	registry.Add("rate/half", Half)
}
