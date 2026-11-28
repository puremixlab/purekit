package toast

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("toast/basic", Basic)
	registry.Add("toast/action", Action)
}
