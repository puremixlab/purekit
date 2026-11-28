package feedback

import (
	"math"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/internal/l10n"
	"github.com/puremixlab/purekit/pkg/theme"
)

type LevelIndicatorProps struct {
	// Label is the accessible name of the level indicator.
	Label string
	// Max sets the upper bound, 100 when zero or negative.
	Max float64
	// Warning and Critical change the status at these values. Nonpositive thresholds are disabled.
	Warning, Critical float64
	// Segments displays discrete steps instead of a continuous fill when greater than one.
	Segments int
}

// LevelIndicator shows a measured level, such as storage or peak volume.
func LevelIndicator(c *ui.Context, value float64, p LevelIndicatorProps) ui.Element {
	tk := theme.TokensOf(c)
	upper := p.Max
	if upper <= 0 || math.IsNaN(upper) || math.IsInf(upper, 0) {
		upper = 100
	}
	if math.IsNaN(value) {
		value = 0
	}
	value = min(max(value, 0), upper)
	fraction := value / upper
	fill, status := tk.Success, ""
	switch {
	case p.Critical > 0 && value >= p.Critical:
		fill, status = tk.Error, l10n.Text(c, "levelCritical")
	case p.Warning > 0 && value >= p.Warning:
		fill, status = tk.Warning, l10n.Text(c, "levelWarning")
	}
	root := ui.Column(c).FillWidth().Role(ui.RoleMeter).Label(p.Label).Range(0, upper, value)
	root.Children(func() {
		track := ui.Row(c).FillWidth().Height(10).Radius(4).Clip().Gap(2)
		track.Children(func() {
			if p.Segments > 1 {
				active := int(math.Round(fraction * float64(p.Segments)))
				for i := range p.Segments {
					color := tk.FillTertiary
					if i < active {
						color = fill
					}
					ui.Box(c.Key(i)).Grow(1).Basis(0).FillHeight().Background(color)
				}
			} else {
				track.Background(tk.FillTertiary)
				ui.Box(c).WidthPercent(float32(fraction * 100)).FillHeight().Background(fill)
			}
		})
		if status != "" {
			ui.Text(c, status).FontSize(12).TextColor(tk.TextSecondary)
		}
	})
	return root
}
