package field

import (
	"strings"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/internal/l10n"
	"github.com/puremixlab/purekit/pkg/lucide"
	"github.com/puremixlab/purekit/pkg/motion"
	"github.com/puremixlab/purekit/pkg/theme"
)

type ComboBoxProps struct {
	// Label is the accessible name of the text field.
	Label string
	// Options are distinct strings that can be chosen from the popup.
	Options []string
	// Placeholder appears while the field is empty.
	Placeholder string
	// Size and Variant set the field's height and surface.
	Size    theme.Size
	Variant InputVariant
	// Disabled prevents typing and opening the options.
	Disabled bool
}

// ComboBox edits value directly and offers a list of choices. Its arrow opens the
// complete list, and typing filters it. Text typed or chosen is kept as the value.
func ComboBox(c *ui.Context, value *string, p ComboBoxProps) ui.Element {
	tk := theme.TokensOf(c)
	root := fieldRow(c, tk, p.Size).Gap(4).Disabled(p.Disabled)
	var input ui.Element
	root.Children(func() {
		cb := ui.ComboboxBase(c, value)
		input = cb.Input.Grow(1).MinWidth(0).FocusRing(false).
			Label(p.Label).Placeholder(p.Placeholder).PlaceholderColor(tk.TextPlaceholder)
		arrow := ui.ButtonBase(c).KeepFocus().Label(l10n.Text(c, "comboBoxOptions")).
			Width(22).Height(22).Center().Radius(tk.BorderRadiusSM).Cursor(ui.CursorPointer)
		if arrow.Hovered() && !p.Disabled {
			arrow.Background(tk.FillTertiary)
		}
		arrow.Children(func() { ui.Icon(c, lucide.ChevronDown()).FontSize(15) })
		if p.Disabled {
			cb.SetOpen(false)
			return
		}
		if arrow.Clicked() {
			cb.SetOpen(!cb.Open())
			input.Focus()
		}
		query := strings.ToLower(strings.TrimSpace(*value))
		filtering := cb.Filtering()
		bounds, anchor := root.Bounds(), input.Bounds()
		cb.Popup(func(panel ui.Element) {
			panel.Margin(5, 0, 0, bounds.X-anchor.X).Padding(4).
				MinWidth(bounds.W).MaxHeight(256).Radius(tk.BorderRadius).
				Background(tk.BgElevated).Transition(motion.Popup()).AttachedOrigin()
			tk.Elevate(panel, theme.BoxShadowSecondary)
			panel.Children(func() {
				matched := false
				for _, option := range p.Options {
					if !filtering || strings.Contains(strings.ToLower(option), query) {
						matched = true
						completeOption(c.Key(option), tk, cb, AutoCompleteOption{Value: option})
					}
				}
				if !matched {
					ui.Text(c, l10n.Text(c, "comboBoxEmpty")).FontSize(13).
						TextColor(tk.TextSecondary).Padding(6, 10)
				}
			})
		})
		if chosen, ok := cb.Chosen(); ok {
			*value = chosen
		}
	})
	styleField(root, tk, p.Variant, false, input.Focused())
	return input
}

type TokenFieldProps struct {
	// Label is the accessible name of the token field.
	Label string
	// Suggestions are offered as the user types, excluding items already in the field.
	Suggestions []string
	// Disabled makes the tokens read-only.
	Disabled bool
}

// TokenField edits a list of strings. Enter or comma accepts an entry and Backspace
// on an empty input removes its last token. Each chip has its own remove button.
func TokenField(c *ui.Context, tokens *[]string, p TokenFieldProps) ui.Element {
	return ui.TokenField(c, tokens, p.Suggestions).Label(p.Label).Disabled(p.Disabled).FillWidth()
}

type StepperProps struct {
	// Label is the accessible name of the stepper.
	Label string
	// Min and Max bound the value. Two zero bounds mean 0 to 100.
	Min, Max float64
	// Step is the increment, 1 when zero or negative.
	Step float64
	// Disabled prevents pointer and keyboard changes.
	Disabled bool
}

// Stepper changes a number with a pair of arrows, independently of a text field.
// The arrow keys work while the stepper has focus.
func Stepper(c *ui.Context, value *float64, p StepperProps) ui.Element {
	lo, hi, step := p.Min, p.Max, p.Step
	if lo == 0 && hi == 0 {
		hi = 100
	}
	if step <= 0 {
		step = 1
	}
	return ui.Stepper(c, value, lo, max(lo, hi), step).Label(p.Label).Disabled(p.Disabled)
}
