package qrcode

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/code"
	"github.com/puremixlab/purekit/pkg/display"
)

func Basic(c *ui.Context) {
	ui.Row(c).Gap(16).AlignItems(ui.Center).Children(func() {
		code.QRCode(c.Key("plain"), "https://github.com/puremixlab/purekit", code.QRCodeProps{})
		code.QRCode(c.Key("icon"), "https://github.com/puremixlab/purekit", code.QRCodeProps{Icon: func(c *ui.Context) {
			display.Avatar(c, "🤯", display.AvatarProps{Size: 36, Background: ui.Hex("#ffffff")})
		}})
		code.QRCode(c.Key("color"), "https://github.com/puremixlab/purekit", code.QRCodeProps{Size: 120,
			Color: ui.Hex("#5b3cc4"), Background: ui.Hex("#f4f0ff"), NoBorder: true})
	})
}
