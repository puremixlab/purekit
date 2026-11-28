package scrollarea

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("scroll-area/basic", Basic)
	registry.Add("scroll-area/horizontal", Horizontal)
}
