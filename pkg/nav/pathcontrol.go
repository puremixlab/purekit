package nav

import (
	"path/filepath"
	"slices"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/action"
	"github.com/puremixlab/purekit/pkg/internal/l10n"
	"github.com/puremixlab/purekit/pkg/lucide"
	"github.com/puremixlab/purekit/pkg/overlay"
	"github.com/puremixlab/purekit/pkg/theme"
)

type PathControlProps struct {
	// Path is a relative or absolute path in the host's file system syntax.
	Path string
	// MaxVisible is the number of path components to show outside the dropdown, 4 when zero.
	MaxVisible int
	// OnNavigate receives the full path of the component selected.
	OnNavigate func(path string)
	// OnDrop receives the component path and any files dropped on it.
	OnDrop func(path string, files []string)
}

type pathComponent struct{ path, label string }

// PathControl shows a file system location with a button for each parent and the current
// location. Long paths put the middle components in a dropdown.
func PathControl(c *ui.Context, p PathControlProps) ui.Element {
	tk := theme.TokensOf(c)
	root := ui.Row(c).Gap(2).AlignItems(ui.Center).Wrap().FillWidth().
		Role(ui.RoleNavigation).Label(l10n.Text(c, "pathControl"))
	segments := pathComponents(p.Path)
	visible := p.MaxVisible
	if visible <= 0 {
		visible = 4
	}
	visible = max(visible, 2)
	open := ui.Local(root, "parents", func() bool { return false })
	root.Children(func() {
		if len(segments) > visible {
			hidden := segments[1 : len(segments)-visible+1]
			pathButton(c.Key(segments[0].path), tk, segments[0], p)
			pathSeparator(c, tk)
			items := make([]overlay.MenuItem, 0, len(hidden))
			for _, component := range hidden {
				items = append(items, overlay.MenuItem{Key: component.path, Label: component.label})
			}
			chosen := overlay.DropdownMenu(c.Key("parents"), open, overlay.DropdownMenuProps{Items: items}, func(b ui.Element) {
				action.StyleButton(c, b.Label(l10n.Text(c, "pathParents")), "…",
					action.ButtonProps{Variant: action.Text, Size: theme.Small})
			})
			if chosen != "" && p.OnNavigate != nil {
				p.OnNavigate(chosen)
			}
			pathSeparator(c, tk)
			segments = segments[len(segments)-visible+1:]
		}
		for i, component := range segments {
			if i > 0 {
				pathSeparator(c, tk)
			}
			pathButton(c.Key(component.path), tk, component, p)
		}
	})
	return root
}

func pathButton(c *ui.Context, tk *theme.Tokens, component pathComponent, p PathControlProps) {
	b := action.Button(c, component.label, action.ButtonProps{Variant: action.Text, Size: theme.Small}).
		Label(component.path)
	if p.OnDrop != nil {
		if b.FileDragOver() {
			b.Background(tk.PrimaryBg)
		}
		if files := b.DroppedFiles(); len(files) > 0 {
			p.OnDrop(component.path, files)
		}
	}
	if p.OnNavigate != nil && b.Clicked() {
		p.OnNavigate(component.path)
	}
}

func pathSeparator(c *ui.Context, tk *theme.Tokens) {
	ui.Icon(c, lucide.ChevronRight()).Size(11, 11).TextColor(tk.TextTertiary)
}

func pathComponents(path string) []pathComponent {
	var components []pathComponent
	if path == "" {
		return components
	}
	for current := filepath.Clean(path); ; current = filepath.Dir(current) {
		parent := filepath.Dir(current)
		label := filepath.Base(current)
		if current == parent || label == "." {
			label = current
		}
		components = append(components, pathComponent{path: current, label: label})
		if current == parent || parent == "." {
			break
		}
	}
	slices.Reverse(components)
	return components
}
