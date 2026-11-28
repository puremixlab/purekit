package modal

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/action"
	"github.com/puremixlab/purekit/pkg/feedback"
	"github.com/puremixlab/purekit/pkg/theme"
)

var deleteOpen bool

func Texts(c *ui.Context) {
	tk := theme.TokensOf(c)
	if action.Button(c, "Delete project", action.ButtonProps{}).Clicked() {
		deleteOpen = true
	}
	p := feedback.ModalProps{Title: "Delete project?", Width: 400, OkText: "Delete", CancelText: "Keep it"}
	if feedback.Modal(c, &deleteOpen, p, func() {
		ui.Text(c, "The project and its 12 files are removed for good.").TextColor(tk.TextSecondary)
	}) {
		deleteOpen = false
	}
}
