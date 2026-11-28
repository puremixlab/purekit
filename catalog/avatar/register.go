package avatar

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("avatar/basic", Basic)
	registry.Add("avatar/fallback", Fallback)
	registry.Add("avatar/sizes", Sizes)
}
