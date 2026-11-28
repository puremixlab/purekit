package empty

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/action"
	"github.com/puremixlab/purekit/pkg/display"
	"github.com/puremixlab/purekit/pkg/lucide"
)

func Basic(c *ui.Context) {
	ui.Column(c).Width(480).Children(func() {
		display.Empty(c, display.EmptyProps{
			Icon:        lucide.Inbox(),
			Title:       "No messages yet",
			Description: "Start a conversation and the messages show here.",
			Action:      func() { action.Button(c, "New chat", action.ButtonProps{Variant: action.Primary}) },
		})
	})
}
