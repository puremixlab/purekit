package checkbox

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("checkbox/basic", Basic)
	registry.Add("checkbox/group", Group)
	registry.Add("checkbox/shapes", Shapes)
}
