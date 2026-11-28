package burger

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("burger/basic", Basic)
}
