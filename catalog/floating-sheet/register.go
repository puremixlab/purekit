package floatingsheet

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("floating-sheet/basic", Basic)
	registry.Add("floating-sheet/inline", Inline)
}
