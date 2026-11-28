package main

import (
	"github.com/puremixlab/purekit/pkg/action"
	"github.com/puremixlab/purekit/pkg/theme"
	"log"
	"os"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	_ "github.com/puremixlab/purekit/catalog/accordion"
	_ "github.com/puremixlab/purekit/catalog/action-icon"
	_ "github.com/puremixlab/purekit/catalog/action-icon-group"
	_ "github.com/puremixlab/purekit/catalog/agent-avatar"
	_ "github.com/puremixlab/purekit/catalog/alert"
	_ "github.com/puremixlab/purekit/catalog/auto-complete"
	_ "github.com/puremixlab/purekit/catalog/avatar"
	_ "github.com/puremixlab/purekit/catalog/badge"
	_ "github.com/puremixlab/purekit/catalog/block"
	_ "github.com/puremixlab/purekit/catalog/breadcrumb"
	_ "github.com/puremixlab/purekit/catalog/burger"
	_ "github.com/puremixlab/purekit/catalog/button"
	_ "github.com/puremixlab/purekit/catalog/checkbox"
	_ "github.com/puremixlab/purekit/catalog/color-picker"
	_ "github.com/puremixlab/purekit/catalog/color-swatches"
	_ "github.com/puremixlab/purekit/catalog/context-menu"
	_ "github.com/puremixlab/purekit/catalog/copy-button"
	_ "github.com/puremixlab/purekit/catalog/date-picker"
	_ "github.com/puremixlab/purekit/catalog/divider"
	_ "github.com/puremixlab/purekit/catalog/draggable-panel"
	_ "github.com/puremixlab/purekit/catalog/draggable-side-nav"
	_ "github.com/puremixlab/purekit/catalog/drawer"
	_ "github.com/puremixlab/purekit/catalog/dropdown-menu"
	_ "github.com/puremixlab/purekit/catalog/editable-text"
	_ "github.com/puremixlab/purekit/catalog/empty"
	_ "github.com/puremixlab/purekit/catalog/floating-panel"
	_ "github.com/puremixlab/purekit/catalog/floating-sheet"
	_ "github.com/puremixlab/purekit/catalog/form"
	_ "github.com/puremixlab/purekit/catalog/grid"
	_ "github.com/puremixlab/purekit/catalog/group-avatar"
	_ "github.com/puremixlab/purekit/catalog/guide-card"
	_ "github.com/puremixlab/purekit/catalog/header"
	_ "github.com/puremixlab/purekit/catalog/highlighter"
	_ "github.com/puremixlab/purekit/catalog/home"
	_ "github.com/puremixlab/purekit/catalog/hotkey"
	_ "github.com/puremixlab/purekit/catalog/hotkey-input"
	_ "github.com/puremixlab/purekit/catalog/icon"
	_ "github.com/puremixlab/purekit/catalog/image"
	_ "github.com/puremixlab/purekit/catalog/image-select"
	_ "github.com/puremixlab/purekit/catalog/input"
	_ "github.com/puremixlab/purekit/catalog/input-number"
	_ "github.com/puremixlab/purekit/catalog/input-otp"
	_ "github.com/puremixlab/purekit/catalog/input-password"
	_ "github.com/puremixlab/purekit/catalog/layout"
	_ "github.com/puremixlab/purekit/catalog/list"
	_ "github.com/puremixlab/purekit/catalog/markdown"
	_ "github.com/puremixlab/purekit/catalog/modal"
	_ "github.com/puremixlab/purekit/catalog/motion"
	_ "github.com/puremixlab/purekit/catalog/outline"
	_ "github.com/puremixlab/purekit/catalog/outline-table"
	_ "github.com/puremixlab/purekit/catalog/pagination"
	_ "github.com/puremixlab/purekit/catalog/popover"
	_ "github.com/puremixlab/purekit/catalog/progress"
	_ "github.com/puremixlab/purekit/catalog/qr-code"
	_ "github.com/puremixlab/purekit/catalog/radio"
	_ "github.com/puremixlab/purekit/catalog/rate"
	"github.com/puremixlab/purekit/catalog/registry"
	_ "github.com/puremixlab/purekit/catalog/result"
	_ "github.com/puremixlab/purekit/catalog/scroll-area"
	_ "github.com/puremixlab/purekit/catalog/search-bar"
	_ "github.com/puremixlab/purekit/catalog/segmented"
	_ "github.com/puremixlab/purekit/catalog/select"
	_ "github.com/puremixlab/purekit/catalog/side-nav"
	_ "github.com/puremixlab/purekit/catalog/skeleton"
	_ "github.com/puremixlab/purekit/catalog/slider"
	_ "github.com/puremixlab/purekit/catalog/snippet"
	_ "github.com/puremixlab/purekit/catalog/sortable-list"
	_ "github.com/puremixlab/purekit/catalog/spin"
	_ "github.com/puremixlab/purekit/catalog/steps"
	_ "github.com/puremixlab/purekit/catalog/switch"
	_ "github.com/puremixlab/purekit/catalog/table"
	_ "github.com/puremixlab/purekit/catalog/tabs"
	_ "github.com/puremixlab/purekit/catalog/tag"
	_ "github.com/puremixlab/purekit/catalog/text"
	_ "github.com/puremixlab/purekit/catalog/text-area"
	_ "github.com/puremixlab/purekit/catalog/toast"
	_ "github.com/puremixlab/purekit/catalog/toggle-group"
	_ "github.com/puremixlab/purekit/catalog/tokens"
	_ "github.com/puremixlab/purekit/catalog/tooltip"
	_ "github.com/puremixlab/purekit/catalog/upload"
	_ "github.com/puremixlab/purekit/catalog/workspace"
)

func main() {
	content := index
	if len(os.Args) > 1 && os.Args[1] != "" {
		view, ok := registry.Get(os.Args[1])
		if !ok {
			log.Fatalf("catalog: no demo named %q; have %v", os.Args[1], registry.Names())
		}
		content = func(c *ui.Context) { preview(c, view) }
	}
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{Title: "PureKit", Width: 720, Height: 480, Content: ui.View(content)})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}

func preview(c *ui.Context, view func(c *ui.Context)) {
	theme.Apply(c)
	ui.Column(c).Fill().Center().Gap(12).Padding(20).
		Background(theme.TokensOf(c).BgContainer).
		Children(func() { view(c) })
}

var chosen string

func index(c *ui.Context) {
	theme.Apply(c)
	names := registry.Names()
	if chosen == "" && len(names) > 0 {
		chosen = names[0]
	}
	tk := theme.TokensOf(c)
	ui.Row(c).Fill().AlignItems(ui.Stretch).Children(func() {
		ui.Scroll(c).Width(200).Background(tk.BgLayout).Children(func() {
			ui.Column(c).Padding(8).Gap(2).Children(func() {
				for _, name := range names {
					variant := action.Text
					if name == chosen {
						variant = action.Filled
					}
					if action.Button(c.Key(name), name, action.ButtonProps{Variant: variant}).AlignSelf(ui.Stretch).Clicked() {
						chosen = name
					}
				}
			})
		})
		ui.Column(c).Grow(1).Children(func() {
			if view, ok := registry.Get(chosen); ok {
				preview(c, view)
			}
		})
	})
}
