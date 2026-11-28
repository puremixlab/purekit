package spin

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("spin/basic", Basic)
	registry.Add("spin/content", Content)
}
