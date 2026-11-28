package badge

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("badge/basic", Basic)
	registry.Add("badge/status", Status)
}
