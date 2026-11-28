package richtexteditor

import (
	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/richtext"
	"github.com/puremixlab/purekit/pkg/field"
)

func Basic(c *ui.Context) {
	root := ui.Column(c).Fill()
	doc := ui.Local(root, "doc", func() *richtext.Document {
		return richtext.NewDocument(richtext.NewText("Release notes\n\nSelect text to change how it looks. Press Command or Ctrl+B to toggle bold."))
	})
	root.Children(func() {
		field.RichTextEditor(c, *doc, field.RichTextEditorProps{Label: "Release note"})
	})
}
