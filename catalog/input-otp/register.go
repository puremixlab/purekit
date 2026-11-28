package inputotp

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("input-otp/basic", Basic)
}
