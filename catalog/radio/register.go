package radio

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("radio/basic", Basic)
	registry.Add("radio/vertical", Vertical)
}
