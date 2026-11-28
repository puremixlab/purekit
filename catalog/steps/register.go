package steps

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("steps/basic", Basic)
	registry.Add("steps/vertical", Vertical)
}
