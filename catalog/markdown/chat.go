package markdown

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/markdown"
	"github.com/puremixlab/purekit/pkg/theme"
)

const reply = "### Sure, here is how\n\nInstall the package, then import it:\n\n" +
	"```sh\ngo get github.com/puremixlab/purekit\n```\n\n" +
	"```go\nimport \"github.com/puremixlab/purekit/pkg/markdown\"\n\nmarkdown.Markdown(c, reply, markdown.MarkdownProps{Variant: markdown.Chat})\n```\n\n" +
	"The **chat** variant uses 14 DIP text, headings barely larger than it, and tighter room between blocks.\n"

func Chat(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Column(c).Width(520).Padding(12, 16).Radius(tk.BorderRadiusLG).Background(tk.FillQuaternary).Children(func() {
		markdown.Markdown(c, reply, markdown.MarkdownProps{Variant: markdown.Chat})
	})
}
