package stepper

import "github.com/puremixlab/purekit/catalog/registry"

func init() { registry.Add("stepper/basic", Basic) }
