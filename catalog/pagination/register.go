package pagination

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("pagination/basic", Basic)
	registry.Add("pagination/options", Options)
}
