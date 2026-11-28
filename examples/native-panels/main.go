package main

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/action"
	"github.com/puremixlab/purekit/pkg/command"
	"github.com/puremixlab/purekit/pkg/desktop"
	"github.com/puremixlab/purekit/pkg/theme"
)

//go:embed status.png
var statusIcon []byte

type app struct {
	window  *mygo.Window
	message string
	font    desktop.Font
	pending bool
}

func main() {
	state := &app{}
	mygo.App.WhenReady(func() {
		state.window = mygo.NewWindow(mygo.WindowOptions{
			Title: "PureKit Native Panels", Width: 610, Height: 330, Content: ui.View(state.view),
		})
		item, err := desktop.NewStatusItem(desktop.StatusItemProps{
			ID: "purekit-native-panels", Title: "PureKit", Tooltip: "PureKit Native Panels", IconPNG: statusIcon,
			Commands: []command.Command{
				{ID: "show", Label: "Show PureKit", Execute: func() { state.window.Show(); state.window.Focus() }},
				{ID: "quit", Label: "Quit", Execute: mygo.App.Quit},
			},
		})
		if err != nil {
			state.message = err.Error()
			state.window.Invalidate()
		} else {
			state.window.OnClosed(item.Tray.Destroy)
		}
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}

func (s *app) view(c *ui.Context) {
	theme.Apply(c)
	tk := theme.TokensOf(c)
	ui.Column(c).Fill().Padding(18).Gap(12).Background(tk.BgContainer).Children(func() {
		ui.Row(c).Gap(6).Wrap().Children(func() {
			if action.Button(c, "Sheet", action.ButtonProps{}).Clicked() {
				s.openWindow(true)
			}
			if action.Button(c, "Floating panel", action.ButtonProps{}).Clicked() {
				s.openWindow(false)
			}
			if action.Button(c, "Choose save location", action.ButtonProps{}).Disabled(s.pending).Clicked() {
				s.chooseSave()
			}
		})
		ui.Row(c).Gap(6).Wrap().Children(func() {
			mac := desktop.SupportsMacPanels()
			share := action.Button(c, "Share", action.ButtonProps{}).Disabled(!mac)
			if share.Clicked() {
				s.report(desktop.ShowShare(s.window, desktop.ShareProps{
					Anchor: share.Bounds(), Text: "Shared from PureKit.",
				}))
			}
			if action.Button(c, "Fonts", action.ButtonProps{}).Disabled(!mac).Clicked() {
				s.report(desktop.ShowFontPanel(desktop.FontPanelProps{Font: s.font, OnChange: func(font desktop.Font) {
					s.window.Update(func() { s.font = font })
				}}))
			}
			if action.Button(c, "Quick Look", action.ButtonProps{}).Disabled(!mac || s.pending).Clicked() {
				s.choosePreview()
			}
		})
		if s.font.Name != "" {
			ui.Text(c, s.font.Family).Font(s.font.Family).FontSize(float32(s.font.Size)).Wrap()
		}
		if s.message != "" {
			ui.Text(c, s.message).TextColor(tk.TextSecondary).Wrap()
		}
	})
}

func (s *app) openWindow(sheet bool) {
	var child *mygo.Window
	content := func(c *ui.Context) {
		ui.Column(c).Fill().Padding(20).Gap(12).Children(func() {
			message := "This panel can stay open while you work."
			if sheet {
				message = "Close the sheet to return to the parent window."
			}
			ui.Text(c, message).Wrap()
			if action.Button(c, "Close", action.ButtonProps{Variant: action.Primary}).Clicked() && child != nil {
				child.Close()
			}
		})
	}
	var err error
	if sheet {
		child, err = desktop.ShowSheet(s.window, desktop.WindowProps{Title: "Native sheet", Width: 350, Height: 145}, content)
	} else {
		child, err = desktop.ShowPanel(desktop.WindowProps{Title: "Floating panel", Width: 350, Height: 145}, content)
	}
	s.report(err)
}

func (s *app) chooseSave() {
	s.pending = true
	go func() {
		path, err := desktop.SaveFile(mygo.SaveDialogOptions{Parent: s.window, Title: "Choose save location"})
		s.window.Update(func() {
			s.pending = false
			s.report(err)
			if err == nil && path != "" {
				s.message = fmt.Sprintf("Save location  %s", path)
			}
		})
	}()
}

func (s *app) choosePreview() {
	s.pending = true
	go func() {
		paths, err := mygo.Dialog.Open(mygo.OpenDialogOptions{Parent: s.window, Title: "Choose a preview file", Multiple: true})
		s.window.Update(func() {
			s.pending = false
			if err == nil && len(paths) > 0 {
				err = desktop.ShowQuickLook(s.window, paths)
			}
			s.report(err)
		})
	}()
}

func (s *app) report(err error) {
	s.message = ""
	if err != nil {
		s.message = err.Error()
	}
}
