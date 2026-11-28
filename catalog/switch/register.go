package switchdemo

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("switch/basic", Basic)
}
