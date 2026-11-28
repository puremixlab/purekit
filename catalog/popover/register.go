package popover

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("popover/basic", Basic)
	registry.Add("popover/placement", Placement)
}
