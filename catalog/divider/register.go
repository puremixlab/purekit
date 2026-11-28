package divider

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("divider/basic", Basic)
	registry.Add("divider/vertical", Vertical)
}
