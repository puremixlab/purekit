package highlighter

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("highlighter/basic", Basic)
	registry.Add("highlighter/variants", Variants)
	registry.Add("highlighter/full-featured", FullFeatured)
}
