package markdown

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("markdown/basic", Basic)
	registry.Add("markdown/chat", Chat)
	registry.Add("markdown/code", Code)
	registry.Add("markdown/stream", Stream)
}
