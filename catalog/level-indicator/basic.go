package levelindicator

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/feedback"
)

func Basic(c *ui.Context) {
	ui.Column(c).Width(285).Gap(7).Children(func() {
		ui.Text(c, "Storage  84 of 100 GB").FontWeight(600)
		feedback.LevelIndicator(c, 84, feedback.LevelIndicatorProps{
			Label: "Storage capacity in GB", Max: 100, Warning: 75, Critical: 95, Segments: 12,
		})
	})
}
