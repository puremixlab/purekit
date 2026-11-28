package layout

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/action"
	"github.com/puremixlab/purekit/pkg/command"
	"github.com/puremixlab/purekit/pkg/display"
	"github.com/puremixlab/purekit/pkg/theme"
)

type ToolbarGroup struct {
	// Commands appear together in the toolbar and its overflow menu.
	Commands []command.Command
}

type ToolbarProps struct {
	// Label is the accessible name of the toolbar.
	Label string
	// Leading and Trailing hold the command groups at either end of the bar.
	Leading, Trailing []ToolbarGroup
}

// Toolbar displays groups of shared commands in a window or a content pane.
// The underlying MyGo toolbar handles arrow-key navigation and narrow-window overflow.
func Toolbar(c *ui.Context, p ToolbarProps) ui.Element {
	tk := theme.TokensOf(c)
	return ui.Toolbar(c, func() {
		toolbarGroups(c, p.Leading, false)
		if len(p.Trailing) > 0 {
			ui.Spacer(c)
			toolbarGroups(c, p.Trailing, true)
		}
	}).Label(p.Label).FillWidth().Padding(6, 8).Gap(10).
		Background(tk.BgContainer).BorderWidth(0, 0, 1, 0).BorderColor(tk.BorderSecondary)
}

type toolbarGroupKey struct {
	trailing bool
	index    int
}

func toolbarGroups(c *ui.Context, groups []ToolbarGroup, trailing bool) {
	for i, group := range groups {
		groupContext := c.Key(toolbarGroupKey{trailing: trailing, index: i})
		ui.Row(groupContext).AlignItems(ui.Center).Gap(3).Children(func() {
			for index, cmd := range group.Commands {
				key := groupContext.Key(index)
				if cmd.ID != "" {
					key = groupContext.Key(cmd.ID)
				}
				toolbarCommand(key, cmd)
			}
		})
	}
}

func toolbarCommand(c *ui.Context, cmd command.Command) {
	enabled, checked := cmd.Available(), cmd.Checked != nil && cmd.Checked()
	var button ui.Element
	if cmd.Checked != nil {
		button = ui.ToggleBase(c, &checked)
	} else {
		button = ui.ButtonBase(c)
	}
	variant := action.Text
	if checked {
		variant = action.Filled
	}
	action.StyleButton(c, button.Disabled(!enabled), "", action.ButtonProps{Size: theme.Small, Variant: variant})
	button.Label(cmd.Label).Children(func() {
		if cmd.Icon != nil {
			ui.Icon(c, cmd.Icon).Size(14, 14)
		}
		ui.Text(c, cmd.Label).SingleLine()
	})
	if !enabled {
		button.Opacity(0.48)
	}
	if mods, key, ok := display.HotkeyBinding(cmd.Shortcut, display.HotkeyAuto); ok {
		button.KeyShortcuts(mods, key)
	}
	if button.Clicked() && cmd.Run() {
		c.Invalidate()
	}
}
