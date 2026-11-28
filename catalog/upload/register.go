package upload

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("upload/basic", Basic)
	registry.Add("upload/dragger", Dragger)
}
