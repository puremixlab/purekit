package alert

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("alert/basic", Basic)
	registry.Add("alert/variants", Variants)
	registry.Add("alert/description", Description)
}
