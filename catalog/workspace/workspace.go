package workspace

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/command"
	"github.com/puremixlab/purekit/pkg/display"
	"github.com/puremixlab/purekit/pkg/feedback"
	"github.com/puremixlab/purekit/pkg/layout"
	"github.com/puremixlab/purekit/pkg/lucide"
	"github.com/puremixlab/purekit/pkg/theme"
)

type document struct {
	name, kind string
	size       int
}

var folders = map[string][]string{
	"Projects": {"Design", "Research"},
	"Archive":  {"2025"},
}

var documents = map[string][]document{
	"Projects": {{"Roadmap.md", "Markdown", 16}, {"Inventory.csv", "Spreadsheet", 32}},
	"Design":   {{"Brand guide.pdf", "PDF", 248}, {"Component notes.md", "Markdown", 56}},
	"Research": {{"Interviews.md", "Markdown", 21}, {"Survey results.csv", "Spreadsheet", 106}},
	"Archive":  {{"Release checklist.md", "Markdown", 19}},
	"2025":     {{"Annual summary.pdf", "PDF", 384}},
}

var columns = []ui.TableColumn{
	{ID: "name", Title: "Name", MinWidth: 128, Sortable: true},
	{ID: "kind", Title: "Kind", Width: 120, Sortable: true},
	{ID: "size", Title: "Size", Width: 69, Align: ui.End, Sortable: true},
}

// State keeps the document browser's selection, layout and commands.
type State struct {
	Layout          layout.SplitViewState
	tree            ui.OutlineState[string]
	treeRow         int
	table           ui.ListState
	selected        int
	sort            ui.SortOrder
	folder          string
	rows            []document
	preview         document
	previewVisible  bool
	commands        []command.Command
	nativeMenuItems []*mygo.MenuItem
}

func New() *State {
	s := &State{folder: "Projects", treeRow: 0, selected: -1}
	s.tree.Open.Add("Projects")
	s.tree.List.Selected = &s.treeRow
	s.tree.List.Label = func(row int) string { return s.tree.Item(row) }
	s.table.Selected = &s.selected
	s.table.Sort = &s.sort
	s.table.Key = func(row int) any { return s.rows[row].name }
	s.table.Label = func(row int) string { return s.rows[row].name }
	s.commands = []command.Command{
		{
			ID: "sidebar", Label: "Sidebar", Icon: lucide.PanelLeft(), Shortcut: "mod+shift+s",
			Checked: func() bool { return !s.Layout.SidebarHidden },
			Execute: func() { s.Layout.SidebarHidden = !s.Layout.SidebarHidden },
		},
		{
			ID: "inspector", Label: "Inspector", Icon: lucide.PanelRightOpen(), Shortcut: "mod+shift+i",
			Checked: func() bool { return !s.Layout.InspectorHidden },
			Execute: func() { s.Layout.InspectorHidden = !s.Layout.InspectorHidden },
		},
		{
			ID: "preview", Label: "Preview", Icon: lucide.Eye(), Shortcut: "mod+shift+p",
			Enabled: func() bool { _, ok := s.selection(); return ok },
			Execute: func() {
				s.preview, s.previewVisible = s.selection()
			},
		},
	}
	return s
}

// NativeMenu gives the native desktop example the same commands as its toolbar.
func (s *State) NativeMenu() *mygo.Menu {
	s.nativeMenuItems = make([]*mygo.MenuItem, len(s.commands))
	for i, cmd := range s.commands {
		s.nativeMenuItems[i] = cmd.MenuItem()
	}
	return mygo.NewMenu([]*mygo.MenuItem{
		{Role: mygo.RoleAppMenu},
		{Role: mygo.RoleEditMenu},
		{Label: "Workspace", Submenu: s.nativeMenuItems},
		{Role: mygo.RoleWindowMenu},
	})
}

func (s *State) View(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Column(c).Fill().Background(tk.BgContainer).Children(func() {
		layout.Toolbar(c, layout.ToolbarProps{
			Label:    "Workspace actions",
			Leading:  []layout.ToolbarGroup{{Commands: s.commands[:2]}},
			Trailing: []layout.ToolbarGroup{{Commands: s.commands[2:]}},
		})
		ui.Box(c).Grow(1).Basis(0).MinHeight(0).Children(func() {
			layout.SplitView(c, &s.Layout, layout.SplitViewProps{
				DefaultSidebarWidth: 158, DefaultDetailPercent: 69,
				Sidebar:   func() { s.sidebar(c) },
				Inspector: func() { s.inspector(c) },
			}, func() { s.fileTable(c) })
		})
		if s.nativeMenuItems == nil {
			command.HandleShortcuts(c, s.commands...)
		}
		for i, item := range s.nativeMenuItems {
			s.commands[i].SyncMenuItem(item)
		}
		feedback.Modal(c, &s.previewVisible, feedback.ModalProps{
			Title: "Document preview", Width: 350, NoFooter: true,
		}, func() {
			ui.Text(c, s.preview.name).FontWeight(600)
			ui.Text(c, fmt.Sprintf("%s  |  %d KB", s.preview.kind, s.preview.size)).TextColor(tk.TextSecondary)
		})
	})
}

func (s *State) sidebar(c *ui.Context) {
	tk := theme.TokensOf(c)
	ui.Column(c).Fill().Padding(9).Gap(6).Background(tk.BgLayout).Children(func() {
		ui.Text(c, "Folders").FontSize(12).FontWeight(600).TextColor(tk.TextSecondary)
		list := display.Outline(c, &s.tree, display.OutlineProps[string]{
			Label: "Folders", Roots: []string{"Projects", "Archive"},
			Children: func(item string) []string { return folders[item] },
		}, func(item string) {
			ui.Icon(c, lucide.Folder()).Size(14, 14)
			ui.Text(c, item).SingleLine()
		}).Grow(1)
		if list.Changed() && s.treeRow >= 0 && s.treeRow < s.tree.Rows() {
			folder := s.tree.Item(s.treeRow)
			if folder != s.folder {
				s.folder = folder
				s.selected = -1
				s.table.ScrollTo(0, ui.Start)
				c.Invalidate()
			}
		}
	})
}

func (s *State) fileTable(c *ui.Context) {
	tk := theme.TokensOf(c)
	s.rows = slices.Clone(documents[s.folder])
	order := s.sort
	slices.SortStableFunc(s.rows, func(a, b document) int {
		var compared int
		switch order.Column {
		case "kind":
			compared = cmp.Compare(a.kind, b.kind)
		case "size":
			compared = cmp.Compare(a.size, b.size)
		default:
			compared = cmp.Compare(a.name, b.name)
		}
		if order.Descending {
			return -compared
		}
		return compared
	})
	ui.Column(c).Fill().MinWidth(0).Padding(9).Gap(6).Children(func() {
		ui.Row(c).Gap(9).AlignItems(ui.Center).Children(func() {
			ui.Text(c, s.folder).FontWeight(600).SingleLine()
			ui.Text(c, fmt.Sprintf("%d files", len(s.rows))).FontSize(12).TextColor(tk.TextSecondary)
		})
		table := display.Table(c, &s.table, display.TableProps{
			Label: "Documents", Columns: columns, Rows: len(s.rows),
		}, func(row, column int) {
			doc := s.rows[row]
			switch column {
			case 0:
				ui.Text(c, doc.name).SingleLine()
			case 1:
				ui.Text(c, doc.kind).SingleLine()
			case 2:
				ui.Text(c, fmt.Sprintf("%d KB", doc.size)).FontFeatures("tnum")
			}
		}).Grow(1)
		if table.Submitted() {
			s.commands[2].Run()
		}
	})
	if s.sort != order {
		c.Invalidate()
	}
}

func (s *State) selection() (document, bool) {
	if s.selected < 0 || s.selected >= len(s.rows) {
		return document{}, false
	}
	return s.rows[s.selected], true
}

func (s *State) inspector(c *ui.Context) {
	tk := theme.TokensOf(c)
	layout.Inspector(c, layout.InspectorProps{
		Title: "Details",
		OnClose: func() {
			s.Layout.InspectorHidden = true
			c.Invalidate()
		},
	}, func() {
		doc, ok := s.selection()
		if !ok {
			ui.Text(c, "Select a document.").TextColor(tk.TextSecondary)
			return
		}
		ui.Icon(c, lucide.FileText()).Size(24, 24).TextColor(tk.Primary)
		ui.Text(c, doc.name).FontWeight(600).Wrap()
		ui.Text(c, doc.kind).TextColor(tk.TextSecondary)
		ui.Text(c, fmt.Sprintf("%d KB", doc.size)).TextColor(tk.TextSecondary)
	})
}
