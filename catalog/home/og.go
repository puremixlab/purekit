package home

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/action"
)

var (
	ogHeadline    = [2]string{"Native UI", "for Go Apps"}
	ogDescription = [2]string{"Components, tokens, and motion", "for desktop apps built with MyGo."}
)

var coverBackground = ui.Hex("#101D1A")
var coverAccent = ui.Hex("#B6F2D0")

// BrandMark draws the mark used by the docs, covers, and catalog showcase.
func BrandMark(c *ui.Context, size float32) {
	ui.Box(c).Size(size, size).Shrink(0).Radius(size * 0.23).Background(ui.Hex("#204338")).Children(func() {
		ui.Box(c).Absolute().Left(size*0.24).Top(size*0.23).
			Size(size*0.17, size*0.54).Radius(size * 0.04).Background(coverAccent)
		ui.Box(c).Absolute().Left(size*0.49).Top(size*0.23).
			Size(size*0.27, size*0.28).Radius(size * 0.07).Background(coverAccent)
	})
}

// OG renders the site's social preview.
func OG(c *ui.Context) {
	ui.Box(c).Fill().Background(coverBackground).Clip().Children(func() {
		ogText(c)
		coverPreview(c)
	})
}

// Banner renders the README header with transparent rounded corners.
func Banner(c *ui.Context) {
	ui.Box(c).Size(1474, 602).Radius(24).Background(coverBackground).Clip().Children(func() {
		ogText(c)
		coverPreview(c)
		ui.Box(c).Absolute().Left(0).Top(0).Right(0).Bottom(0).Radius(24).
			Border(1, ui.RGBA(255, 255, 255, 0.12)).IgnoreInput(true)
	})
}

func ogText(c *ui.Context) {
	white, grey := ui.Hex("#F4FCF7"), ui.Hex("#A8BFB5")
	ui.Row(c).Absolute().Left(64).Top(104).Gap(16).AlignItems(ui.Center).Children(func() {
		BrandMark(c, 56)
		ui.Text(c, "PureKit").FontSize(34).FontWeight(700).TextColor(white)
	})
	const headlineWidth = 680 - 64
	font := ui.Font{Family: "system-ui", Size: 76, Weight: 700}
	for _, line := range ogHeadline {
		width, _ := ui.ShapeText(line, font, 0).Size()
		if width > headlineWidth {
			font.Size *= headlineWidth / width
		}
	}
	ui.Column(c).Absolute().Left(64).Top(210).Children(func() {
		for i, line := range ogHeadline {
			ui.Text(c.Key(i), line).Font(font.Family).FontSize(font.Size).
				FontWeight(font.Weight).LetterSpacing(-1.5).LineHeight(1.12).TextColor(white)
		}
	})
	ui.Column(c).Absolute().Left(64).Top(458).Children(func() {
		for i, line := range ogDescription {
			ui.Text(c.Key(i), line).FontSize(24).LineHeight(1.4).TextColor(grey)
		}
	})
}

func coverPreview(c *ui.Context) {
	border := ui.Hex("#3A5E51")
	ui.Column(c).Absolute().Right(64).Top(138).Width(446).Padding(26).Gap(20).
		Radius(26).Border(1, border).Background(ui.Hex("#1A3028")).Children(func() {
		ui.Row(c).AlignItems(ui.Center).Gap(12).Children(func() {
			BrandMark(c, 32)
			ui.Text(c, "Components").FontSize(24).FontWeight(650).TextColor(ui.Hex("#F4FCF7"))
		})
		ui.Column(c).Gap(14).Padding(18).Radius(16).Border(1, border).
			Background(coverBackground).Children(func() {
			ui.Text(c, "Design tokens").FontSize(17).TextColor(ui.Hex("#C1D8CF"))
			ui.Row(c).Gap(12).Children(func() {
				for i, color := range []string{"#B6F2D0", "#7DCDB2", "#4F8F7A", "#31594B", "#243E34"} {
					ui.Box(c.Key(i)).Size(51, 38).Radius(8).Background(ui.Hex(color))
				}
			})
		})
		ui.Row(c).Gap(12).Children(func() {
			action.Button(c, "Primary", action.ButtonProps{Variant: action.Primary}).
				Background(coverAccent).TextColor(coverBackground)
			action.Button(c, "Outlined", action.ButtonProps{}).Border(1, border).TextColor(ui.Hex("#F4FCF7"))
		})
	})
}
