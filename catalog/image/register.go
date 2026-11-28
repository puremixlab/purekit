package image

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("image/basic", Basic)
	registry.Add("image/states", States)
	registry.Add("image/gallery", Gallery)
}
