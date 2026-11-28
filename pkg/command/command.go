// Package command shares an action between a toolbar, shortcuts and a native menu.
package command

import (
	"strings"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/display"
)

type Command struct {
	// ID is a stable identifier, unique among the commands in a menu or toolbar.
	ID string
	// Label names the action in buttons and menus.
	Label string
	// Icon is optional artwork for the toolbar button.
	Icon *ui.SVG
	// Shortcut uses Hotkey syntax, such as "mod+shift+i".
	Shortcut string
	// Enabled reports the current availability. Nil enables the action when Execute is set.
	Enabled func() bool
	// Checked makes the command a toggle and reports its current value.
	Checked func() bool
	// Execute performs the action, including updating the value of a toggle.
	Execute func()
}

// Available reports whether the command can run.
func (cmd Command) Available() bool {
	return cmd.Execute != nil && (cmd.Enabled == nil || cmd.Enabled())
}

// Run invokes the action only when it is currently available.
func (cmd Command) Run() bool {
	if !cmd.Available() {
		return false
	}
	cmd.Execute()
	return true
}

// HandleShortcuts runs the first enabled command that matches the current key event.
// Use it in views where native menu accelerators aren't already handling these commands.
func HandleShortcuts(c *ui.Context, commands ...Command) bool {
	for _, cmd := range commands {
		if !cmd.Available() {
			continue
		}
		mods, key, ok := display.HotkeyBinding(cmd.Shortcut, display.HotkeyAuto)
		if ok && c.Shortcut(mods, key) {
			return cmd.Run()
		}
	}
	return false
}

// MenuItem builds a MyGo native menu item with the same action and shortcut.
func (cmd Command) MenuItem() *mygo.MenuItem {
	item := &mygo.MenuItem{
		ID: cmd.ID, Label: cmd.Label, Accelerator: nativeShortcut(cmd.Shortcut), Disabled: !cmd.Available(),
		Click: func(item *mygo.MenuItem, _ *mygo.Window) {
			cmd.Run()
			cmd.SyncMenuItem(item)
		},
	}
	if cmd.Checked != nil {
		item.Type = mygo.MenuItemCheckbox
		item.Checked = cmd.Checked()
	}
	return item
}

// SyncMenuItem updates a menu item after the app changes the command's availability or value.
func (cmd Command) SyncMenuItem(item *mygo.MenuItem) {
	if item == nil {
		return
	}
	if enabled := cmd.Available(); item.IsEnabled() != enabled {
		item.SetEnabled(enabled)
	}
	if cmd.Checked != nil {
		if checked := cmd.Checked(); item.IsChecked() != checked {
			item.SetChecked(checked)
		}
	}
}

func nativeShortcut(shortcut string) string {
	parts := strings.Split(shortcut, "+")
	for i, part := range parts {
		part = strings.TrimSpace(part)
		if strings.EqualFold(part, "mod") {
			part = "CmdOrCtrl"
		}
		parts[i] = part
	}
	return strings.Join(parts, "+")
}
