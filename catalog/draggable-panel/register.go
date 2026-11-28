package draggablepanel

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("draggable-panel/basic", Basic)
	registry.Add("draggable-panel/placement", Placement)
}
