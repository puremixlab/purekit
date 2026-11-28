// Package desktop opens native windows, panels and menus from PureKit desktop apps.
package desktop

import (
	"errors"
	"runtime"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/command"
	"github.com/puremixlab/purekit/pkg/theme"
)

var (
	// ErrUnavailable means the current platform has no implementation of the native control.
	ErrUnavailable = errors.New("native control is unavailable on this platform")
	// ErrNotReady means the MyGo app has not entered its event loop.
	ErrNotReady = errors.New("the desktop app is not ready")
	// ErrWindow means the action needs an open parent window.
	ErrWindow = errors.New("an open parent window is required")
)

type WindowProps struct {
	// Title is the window's title.
	Title string
	// Width and Height are the size in DIPs, defaulting to 460 by 340.
	Width, Height int
	// Theme is the theme configuration for the content. Nil uses PureKit's defaults.
	Theme *theme.Config
}

// ShowSheet opens a native child window that blocks input to its parent until it closes.
// On macOS this is a window-attached sheet.
func ShowSheet(parent *mygo.Window, p WindowProps, content func(*ui.Context)) (*mygo.Window, error) {
	if err := nativeReady(); err != nil {
		return nil, err
	}
	if parent == nil || parent.NativeHandle() == 0 {
		return nil, ErrWindow
	}
	opts := windowOptions(p, content)
	opts.Parent, opts.Modal = parent, true
	return mygo.NewWindow(opts), nil
}

// ShowPanel opens an independent floating panel with PureKit content.
func ShowPanel(p WindowProps, content func(*ui.Context)) (*mygo.Window, error) {
	if err := nativeReady(); err != nil {
		return nil, err
	}
	opts := windowOptions(p, content)
	opts.Panel = true
	return mygo.NewWindow(opts), nil
}

func windowOptions(p WindowProps, content func(*ui.Context)) mygo.WindowOptions {
	width, height := p.Width, p.Height
	if width <= 0 {
		width = 460
	}
	if height <= 0 {
		height = 340
	}
	config := p.Theme
	if config == nil {
		config = theme.Default()
	}
	return mygo.WindowOptions{
		Title: p.Title, Width: width, Height: height,
		Content: ui.View(func(c *ui.Context) {
			config.Apply(c)
			content(c)
		}),
	}
}

// SaveFile asks the system for a save location. It returns an empty path on cancellation
// and leaves writing the file to the caller.
func SaveFile(p mygo.SaveDialogOptions) (string, error) {
	if err := nativeReady(); err != nil {
		return "", err
	}
	return mygo.Dialog.Save(p)
}

type StatusItemProps struct {
	// ID identifies the item across launches.
	ID string
	// Title and Tooltip describe the item to the system.
	Title, Tooltip string
	// IconPNG is the icon image. On macOS a template icon follows the menu bar appearance.
	IconPNG        []byte
	IconIsTemplate bool
	// Commands are the actions shown in the item's menu.
	Commands []command.Command
}

// StatusItem is an app menu bar item on macOS, or an icon in the notification area elsewhere.
type StatusItem struct {
	Tray     *mygo.Tray
	commands []command.Command
	items    []*mygo.MenuItem
}

// NewStatusItem creates a status item whose menu uses the app's shared commands.
func NewStatusItem(p StatusItemProps) (*StatusItem, error) {
	if err := nativeReady(); err != nil {
		return nil, err
	}
	items := make([]*mygo.MenuItem, len(p.Commands))
	for i, cmd := range p.Commands {
		items[i] = cmd.MenuItem()
	}
	tray, err := mygo.NewTray(mygo.TrayOptions{
		ID: p.ID, Title: p.Title, ToolTip: p.Tooltip, Icon: p.IconPNG,
		IconIsTemplate: p.IconIsTemplate, Menu: mygo.NewMenu(items),
	})
	if err != nil {
		return nil, err
	}
	return &StatusItem{Tray: tray, commands: append([]command.Command(nil), p.Commands...), items: items}, nil
}

// Refresh updates the availability and checked state of the status item's commands.
func (item *StatusItem) Refresh() {
	for i, cmd := range item.commands {
		cmd.SyncMenuItem(item.items[i])
	}
}

func nativeReady() error {
	if runtime.GOOS == "js" {
		return ErrUnavailable
	}
	if !mygo.App.IsReady() {
		return ErrNotReady
	}
	return nil
}
