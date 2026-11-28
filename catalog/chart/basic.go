package chart

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/display"
)

func Basic(c *ui.Context) {
	display.Chart(c, display.ChartProps{
		Label: "Document activity", Description: "Documents updated by weekday",
		Points: []display.ChartPoint{
			{Label: "Mon", Value: 16}, {Label: "Tue", Value: 22}, {Label: "Wed", Value: 13},
			{Label: "Thu", Value: 30}, {Label: "Fri", Value: 24},
		},
		Height: 173,
	})
}
