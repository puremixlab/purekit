package autocomplete

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("auto-complete/basic", Basic)
}
