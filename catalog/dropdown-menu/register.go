package dropdownmenu

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("dropdown-menu/basic", Basic)
	registry.Add("dropdown-menu/action-icon", ActionIconTrigger)
}
