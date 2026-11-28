package field

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/action"
	"github.com/puremixlab/purekit/pkg/display"
	"github.com/puremixlab/purekit/pkg/internal/l10n"
	"github.com/puremixlab/purekit/pkg/lucide"
	"github.com/puremixlab/purekit/pkg/theme"
)

// ImageWellValue holds the displayed image and the file it was picked from, if any.
type ImageWellValue struct {
	Image ui.ImageSource
	File  File
}

type ImageWellProps struct {
	// Label names the image well and the selection button.
	Label string
	// Width and Height size the preview, defaulting to 210 by 132 DIPs.
	Width, Height float32
	// Disabled prevents replacing or removing the displayed image.
	Disabled bool
	// Picker is an optional replacement for the system or browser file chooser.
	Picker func(PickOptions) ([]File, error)
}

// ImageWell accepts an image by drop or from a file picker, previews it and lets people clear it.
// A failed decode leaves the existing image displayed.
func ImageWell(c *ui.Context, value *ImageWellValue, p ImageWellProps) ui.Element {
	tk := theme.TokensOf(c)
	width, height := p.Width, p.Height
	if width <= 0 {
		width = 210
	}
	if height <= 0 {
		height = 132
	}
	label := p.Label
	if label == "" {
		label = l10n.Text(c, "imageWellChoose")
	}
	root := ui.Column(c).Width(width).Gap(7).Role(ui.RoleGroup).Label(label)
	failed := ui.Local(root, "failed", func() bool { return false })
	set := func(file File) {
		data, err := file.ReadAll()
		if err == nil {
			var image ui.ImageSource
			image, err = ui.DecodeImage(data)
			if err == nil {
				*value = ImageWellValue{Image: image, File: file}
			}
		}
		*failed = err != nil
		c.Invalidate()
	}
	root.Children(func() {
		frame := ui.Box(c).Size(width, height).Center().Radius(tk.BorderRadius).
			Border(1, tk.BorderSecondary).Background(tk.FillQuaternary)
		if frame.FileDragOver() && !p.Disabled {
			frame.Border(1, tk.Primary).Background(tk.PrimaryBg)
		}
		if files := frame.DroppedFiles(); len(files) > 0 && !p.Disabled {
			for _, path := range files {
				file := fileAt(path)
				if matchesAccept(file, "image/*") {
					set(file)
					break
				}
			}
		}
		frame.Children(func() {
			if value.Image != nil {
				display.Image(c, value.Image, display.ImageProps{
					Width: width, Height: height, Contain: true, Variant: display.ImageBorderless,
					Alt: value.File.Name,
				})
				return
			}
			ui.Column(c).Center().Gap(8).Children(func() {
				ui.Icon(c, lucide.Image()).Size(24, 24).TextColor(tk.TextSecondary)
				ui.Text(c, l10n.Text(c, "imageWellEmpty")).FontSize(12).TextColor(tk.TextSecondary)
			})
		})
		ui.Row(c).Gap(6).Children(func() {
			trigger := Upload(c, UploadProps{
				Accept: "image/*", Picker: p.Picker, Disabled: p.Disabled,
				OnFiles: func(files []File) { set(files[0]) },
			}, nil)
			action.StyleButton(c, trigger.Label(label), label, action.ButtonProps{Size: theme.Small})
			if value.Image != nil && action.Button(c, l10n.Text(c, "clear"), action.ButtonProps{
				Size: theme.Small, Variant: action.Text,
			}).Disabled(p.Disabled).Clicked() && !p.Disabled {
				*value = ImageWellValue{}
				*failed = false
				c.Invalidate()
			}
		})
		if *failed {
			ui.Text(c, l10n.Text(c, "imageWellInvalid")).Role(ui.RoleAlert).
				FontSize(12).TextColor(tk.Error)
		}
	})
	return root
}
