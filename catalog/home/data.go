package home

import "github.com/puremixlab/purekit/pkg/display"

type message struct {
	user bool
	text string
}

type conversation struct {
	name, topic string
	shape       display.AgentShape
	color       display.AgentColor
	messages    []message
}

const firstReply = "Every component on this page is drawn on the GPU by **MyGo**, straight from Go: no webview, no HTML, no JavaScript.\n\n" +
	"- One small binary per platform, cross-compiled without cgo\n" +
	"- Shared tokens for color, radius, and typography\n" +
	"- This page runs that same Go code through WebAssembly\n\n" +
	"```go\nfunc Hello(c *ui.Context) {\n\ttheme.Apply(c)\n\taction.Button(c, \"Hello\", action.ButtonProps{Variant: action.Primary})\n}\n```\n\n" +
	"Drag the side nav's edge, collapse it, or send a message. That is the whole app."

var conversations = []conversation{
	{name: "PureKit Assistant", topic: "How MyGo draws", shape: display.AgentCloud, color: display.AgentPurple, messages: []message{
		{user: true, text: "How does PureKit draw its components?"},
		{text: firstReply},
	}},
	{name: "Go Reviewer", topic: "Review a pull request", shape: display.AgentHex, color: display.AgentGeekblue, messages: []message{
		{user: true, text: "Can you review the upload queue fix?"},
		{text: "The race is gone: the queue now holds its lock while it **swaps the batch**, and the new test fails without the fix.\n\nOne nit: `flush` can return early when the batch is empty."},
	}},
	{name: "Designer", topic: "Pick a palette", shape: display.AgentCat, color: display.AgentMagenta, messages: []message{
		{user: true, text: "Which accent should the release page use?"},
		{text: "Geekblue for links and the primary button, with **volcano** for the one call to action that matters. Both are in the theme palette, so a web page and a native app can share them."},
	}},
	{name: "Translator", topic: "Release notes in Japanese", shape: display.AgentGhost, color: display.AgentOrange, messages: []message{
		{user: true, text: "Translate: \"Native UI, drawn in Go.\""},
		{text: "ネイティブ UI を、Go で描画。"},
	}},
	{name: "Writer", topic: "Launch post", shape: display.AgentStar, color: display.AgentGray, messages: []message{
		{user: true, text: "One line for the launch post?"},
		{text: "PureKit brings reusable UI components to native Go desktop apps."},
	}},
}

var replies = []string{
	"Sent from a Go text area, answered by a Go Markdown view, streamed a character at a time. Nothing here is HTML.",
	"Try the light and dark switch in the docs header: MyGo follows it live, with the same tokens as the React components.",
	"Every list row, avatar and button you see is a PureKit component. The catalog has the code for each.",
}
