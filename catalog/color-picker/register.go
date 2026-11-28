package colorpicker

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("color-picker/basic", Basic)
	registry.Add("color-picker/triggers", Triggers)
}
