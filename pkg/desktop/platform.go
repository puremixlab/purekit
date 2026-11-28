package desktop

import (
	"runtime"

	"github.com/egoist/mygo/ui"
)

// SupportsMacPanels reports whether the caller can use the AppKit-only panels.
func SupportsMacPanels() bool { return runtime.GOOS == "darwin" }

type ShareProps struct {
	// Anchor is a rectangle in the parent window's content coordinates.
	Anchor ui.Rect
	// Text and Files are offered to the system's sharing services.
	Text  string
	Files []string
}

// Font describes a typeface selected in the system font panel.
type Font struct {
	// Name is the PostScript font name. Family is the user-facing family name.
	Name, Family string
	// Size is the font size in points.
	Size float64
}

type FontPanelProps struct {
	// Font is the initial selection. A zero Font uses the system font.
	Font Font
	// OnChange receives choices from the system font panel on the app's UI thread.
	OnChange func(Font)
}
