package layout

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/action"
	"github.com/puremixlab/purekit/pkg/internal/l10n"
	"github.com/puremixlab/purekit/pkg/lucide"
	"github.com/puremixlab/purekit/pkg/theme"
)

// SplitViewState holds the layout values an app can save and restore between windows.
type SplitViewState struct {
	// SidebarWidth is the leading pane width in DIPs.
	SidebarWidth float32
	// DetailPercent is the percentage of the space after the sidebar given to the main content.
	DetailPercent float32
	// SidebarHidden and InspectorHidden hide their panes and preserve their sizes.
	SidebarHidden, InspectorHidden bool
}

type SplitViewProps struct {
	// Sidebar and Inspector build the optional leading and trailing panes.
	Sidebar, Inspector func()
	// DefaultSidebarWidth is the initial and reset width in DIPs, 200 when zero.
	DefaultSidebarWidth float32
	// DefaultDetailPercent is the initial and reset share of the main pane, 68 when zero.
	DefaultDetailPercent float32
}

// SplitView lays out a sidebar, the main view and an optional inspector.
// Drag a divider, or focus it and use the arrow keys, to resize its panes.
// Double-click a divider to restore its default size.
func SplitView(c *ui.Context, state *SplitViewState, p SplitViewProps, content func()) ui.Element {
	sidebarWidth := p.DefaultSidebarWidth
	if sidebarWidth <= 0 {
		sidebarWidth = 200
	}
	detailPercent := p.DefaultDetailPercent
	if detailPercent <= 0 || detailPercent >= 100 {
		detailPercent = 68
	}
	if state.SidebarWidth <= 0 {
		state.SidebarWidth = sidebarWidth
	}
	if state.DetailPercent <= 0 || state.DetailPercent >= 100 {
		state.DetailPercent = detailPercent
	}
	main := func() {
		if p.Inspector == nil || state.InspectorHidden {
			content()
			return
		}
		ui.SplitPercent(c, &state.DetailPercent, content, p.Inspector).
			Fill().MinWidth(0).MinHeight(0).SplitReset(detailPercent)
	}
	return ui.Box(c).Fill().MinWidth(0).MinHeight(0).Children(func() {
		if p.Sidebar == nil || state.SidebarHidden {
			main()
			return
		}
		ui.Split(c, &state.SidebarWidth, p.Sidebar, main).
			Fill().MinWidth(0).MinHeight(0).SplitReset(sidebarWidth)
	})
}

type InspectorProps struct {
	// Title is displayed above the contents and names the inspector for assistive technology.
	Title string
	// OnClose adds a close button when set.
	OnClose func()
}

// Inspector builds a labeled, scrolling details pane for the current selection.
func Inspector(c *ui.Context, p InspectorProps, content func()) ui.Element {
	tk := theme.TokensOf(c)
	root := ui.Column(c).Fill().MinWidth(0).MinHeight(0).
		Role(ui.RoleRegion).Label(p.Title).Background(tk.BgLayout)
	root.Children(func() {
		ui.Row(c).AlignItems(ui.Center).Gap(8).Padding(8, 12).
			BorderWidth(0, 0, 1, 0).BorderColor(tk.BorderSecondary).Children(func() {
			ui.Text(c, p.Title).FontWeight(600).SingleLine().Grow(1).MinWidth(0)
			if p.OnClose != nil && action.ActionIcon(c, lucide.PanelRightClose(), action.ActionIconProps{
				Title: l10n.Text(c, "close"), Size: theme.Small,
			}).Clicked() {
				p.OnClose()
			}
		})
		ui.Scroll(c).Grow(1).Basis(0).MinHeight(0).Children(func() {
			ui.Column(c).FillWidth().Padding(12).Gap(8).Children(content)
		})
	})
	return root
}
