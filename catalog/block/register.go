package block

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("block/basic", Basic)
}
