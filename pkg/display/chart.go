package display

import (
	"math"
	"strconv"

	"github.com/egoist/mygo/ui"
	"github.com/egoist/mygo/ui/chart"
	"github.com/puremixlab/purekit/pkg/theme"
)

type ChartKind int

const (
	ChartBars ChartKind = iota
	ChartLine
)

type ChartPoint struct {
	Label string
	Value float64
}

type ChartProps struct {
	// Label and Description name and summarize the data for assistive technology.
	Label, Description string
	// Points are the categories and values in display order.
	Points []ChartPoint
	// Kind selects bars or a line.
	Kind ChartKind
	// Height is the plot height in DIPs, 176 when zero.
	Height float32
	// Selected tracks the highlighted point by index, or -1 for no point.
	Selected *int
	// FormatValue formats the value shown when a point is highlighted.
	FormatValue func(float64) string
}

// Chart presents a categorical data series with axes and pointer and keyboard selection.
func Chart(c *ui.Context, p ChartProps) ui.Element {
	tk := theme.TokensOf(c)
	root := ui.Column(c).FillWidth().Gap(4)
	index := p.Selected
	if index == nil {
		index = ui.Local(root, "point", func() int { return -1 })
	}
	height := p.Height
	if height <= 0 {
		height = 176
	}
	labels := make([]string, len(p.Points))
	values := make([]float64, len(p.Points))
	low, high := 0.0, 0.0
	for i, point := range p.Points {
		labels[i] = point.Label
		if !math.IsNaN(point.Value) && !math.IsInf(point.Value, 0) {
			values[i] = point.Value
			low, high = min(low, point.Value), max(high, point.Value)
		}
	}
	if low == high {
		high = low + 1
	}
	x := chart.NewBand(labels...).Padding(0.3, 0.12)
	y := chart.NewLinear(low, high).Nice(4)
	xs, ys := make([]float32, len(values)), make([]float32, len(values))
	for i, value := range values {
		xs[i], ys[i] = x.At(i), y.At(value)
	}
	describe := func(i int) string {
		value := strconv.FormatFloat(values[i], 'f', -1, 64)
		if p.FormatValue != nil {
			value = p.FormatValue(values[i])
		}
		return labels[i] + "  " + value
	}
	root.Children(func() {
		xTicks := x.Ticks()
		if len(xTicks) > 8 {
			stride := (len(xTicks) + 7) / 8
			shown := make([]chart.Tick, 0, 8)
			for i := 0; i < len(xTicks); i += stride {
				shown = append(shown, xTicks[i])
			}
			xTicks = shown
		}
		yTicks := y.Ticks(4)
		chart.Plot(c, chart.Axes{Left: yTicks, Bottom: xTicks}, func(area ui.Element) {
			area.Draw(func(painter *ui.Painter, rect ui.Rect) {
				for _, tick := range yTicks {
					row := chart.Y(rect, tick.At)
					painter.Line(rect.X, row, rect.X+rect.W, row, 1, tk.BorderSecondary)
				}
				if p.Kind == ChartLine && len(xs) > 1 {
					painter.StrokePath(chart.Line(rect, xs, ys, chart.CurveLinear), 2, tk.Primary)
				}
			})
			chart.Select(c, area, index, xs, describe)
			for i, value := range values {
				if p.Kind == ChartBars {
					color := tk.Primary.Alpha(0.72)
					if *index == i {
						color = tk.Primary
					}
					chart.Rect(c.Key(i), x.Band(i), y.Span(0, value)).Background(color).Radius(3)
				} else if len(values) == 1 || *index == i {
					chart.Point(c.Key(i), xs[i], ys[i], 8, 8).Radius(4).Background(tk.Primary)
				}
			}
		}).Height(height).Gap(7).Padding(8, 12).Label(p.Label).Description(p.Description).
			FontSize(12).TextColor(tk.TextSecondary)
		ui.Box(c).Height(19).Children(func() {
			if i := *index; i >= 0 && i < len(values) {
				ui.Text(c, describe(i)).SingleLine().FontSize(12).TextColor(tk.Text)
			}
		})
	})
	return root
}
