//go:build !darwin

package desktop

import "github.com/egoist/mygo"

// ShowShare opens the system share picker on macOS.
func ShowShare(_ *mygo.Window, _ ShareProps) error { return ErrUnavailable }

// ShowFontPanel opens the system font panel on macOS.
func ShowFontPanel(_ FontPanelProps) error { return ErrUnavailable }

// ShowQuickLook previews local files through Quick Look on macOS.
func ShowQuickLook(_ *mygo.Window, _ []string) error { return ErrUnavailable }
