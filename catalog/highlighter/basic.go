package highlighter

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/code"
)

func Basic(c *ui.Context) {
	ui.Column(c).Gap(16).Width(560).Children(func() {
		code.Highlighter(c.Key("go"), goSample, code.HighlighterProps{Language: "go"})
		code.Highlighter(c.Key("ts"), tsSample, code.HighlighterProps{Language: "ts"})
	})
}
