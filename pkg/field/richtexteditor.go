package field

import (
	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/richtext"
	"github.com/puremixlab/purekit/pkg/command"
	"github.com/puremixlab/purekit/pkg/internal/l10n"
	"github.com/puremixlab/purekit/pkg/layout"
	"github.com/puremixlab/purekit/pkg/lucide"
)

type RichTextEditorProps struct {
	// Label is the accessible name of the editor.
	Label string
	// Placeholder is visible when the document is empty.
	Placeholder string
	// ReadOnly keeps the document selectable without permitting edits.
	ReadOnly bool
	// HideToolbar shows just the editable document.
	HideToolbar bool
	// OnChange is called with the document's text after typing or a toolbar edit.
	OnChange func(richtext.Text)
}

// RichTextEditor edits a MyGo rich text document with a formatting toolbar.
// Keep the document for the lifetime of the editing session.
func RichTextEditor(c *ui.Context, doc *richtext.Document, p RichTextEditorProps) ui.Element {
	root := ui.Column(c).Fill().MinWidth(0).MinHeight(0)
	root.Children(func() {
		version := doc.Version()
		if !p.ReadOnly && !p.HideToolbar {
			formatting := doc.Formatting()
			layout.Toolbar(c, layout.ToolbarProps{
				Label: l10n.Text(c, "richTextFormat"), KeepFocus: true,
				Leading: []layout.ToolbarGroup{{Commands: []command.Command{
					{ID: "bold", Label: l10n.Text(c, "richTextBold"), Icon: lucide.Bold(), Shortcut: "mod+b",
						Checked: func() bool { return formatting.Bold == richtext.FormatOn }, Execute: doc.ToggleBold},
					{ID: "italic", Label: l10n.Text(c, "richTextItalic"), Icon: lucide.Italic(), Shortcut: "mod+i",
						Checked: func() bool { return formatting.Italic == richtext.FormatOn }, Execute: doc.ToggleItalic},
					{ID: "underline", Label: l10n.Text(c, "richTextUnderline"), Icon: lucide.Underline(), Shortcut: "mod+u",
						Checked: func() bool { return formatting.Underline == richtext.FormatOn }, Execute: doc.ToggleUnderline},
				}}, {Commands: []command.Command{
					{ID: "undo", Label: l10n.Text(c, "richTextUndo"), Icon: lucide.Undo2(), Enabled: doc.CanUndo, Execute: doc.Undo},
					{ID: "redo", Label: l10n.Text(c, "richTextRedo"), Icon: lucide.Redo2(), Enabled: doc.CanRedo, Execute: doc.Redo},
				}}},
			})
		}
		editor := richtext.Editor(c, doc, richtext.EditorOptions{
			Label: p.Label, Placeholder: p.Placeholder, ReadOnly: p.ReadOnly,
		}).Grow(1).Basis(0).MinHeight(0)
		if p.OnChange != nil && (editor.Changed() || doc.Version() != version) {
			p.OnChange(doc.Text())
		}
	})
	return root
}
