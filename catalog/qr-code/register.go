package qrcode

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("qr-code/basic", Basic)
	registry.Add("qr-code/status", Status)
}
