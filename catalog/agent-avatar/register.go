package agentavatar

import "github.com/puremixlab/purekit/catalog/registry"

func init() {
	registry.Add("agent-avatar/basic", Basic)
	registry.Add("agent-avatar/states", States)
}
