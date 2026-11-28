package empty

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("empty/basic", Basic)
	registry.Add("empty/page", Page)
}
