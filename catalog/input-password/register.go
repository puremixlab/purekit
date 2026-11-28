package inputpassword

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("input-password/basic", Basic)
}
