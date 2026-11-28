package drawer

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("drawer/basic", Basic)
	registry.Add("drawer/placement", Placement)
}
