package tokens

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("tokens/swatches", Swatches)
}
