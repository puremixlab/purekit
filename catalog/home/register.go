package home

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("home/showcase", Showcase)
}
