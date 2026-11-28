package tooltip

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("tooltip/basic", Basic)
	registry.Add("tooltip/placement", Placement)
}
