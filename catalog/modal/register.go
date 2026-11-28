package modal

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("modal/basic", Basic)
	registry.Add("modal/texts", Texts)
}
