package icon

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("icon/sizes", Sizes)
	registry.Add("icon/spin", Spin)
}
