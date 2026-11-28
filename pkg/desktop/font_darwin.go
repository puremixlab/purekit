package desktop

import (
	"math"
	"sync"

	"github.com/ebitengine/purego/objc"
)

type fontSelection struct {
	manager, controller, font, previousTarget objc.ID
	previousAction                            objc.SEL
	onChange                                  func(Font)
}

var activeFontPanel *fontSelection

var fontControllerClass = sync.OnceValues(func() (objc.Class, error) {
	return objc.RegisterClass("PureKitFontPanelController", objc.GetClass("NSObject"), nil, nil, []objc.MethodDef{
		{Cmd: selector("pureKitChangeFont:"), Fn: fontChanged},
		{Cmd: selector("pureKitFontPanelClosed:"), Fn: fontPanelClosed},
	})
})

// ShowFontPanel shows the macOS Font panel and sends the chosen face and size to OnChange.
func ShowFontPanel(p FontPanelProps) error {
	return mainThread(func() error {
		controllerClass, err := fontControllerClass()
		if err != nil {
			return err
		}
		manager := send(class("NSFontManager"), "sharedFontManager")
		panel := send(manager, "fontPanel:", true)
		if panel == 0 {
			return ErrUnavailable
		}
		finishFontPanel()
		size := p.Font.Size
		if size <= 0 || math.IsNaN(size) || math.IsInf(size, 0) {
			size = objc.Send[float64](class("NSFont"), selector("systemFontSize"))
		}
		var font objc.ID
		if p.Font.Name != "" {
			font = send(class("NSFont"), "fontWithName:size:", nativeString(p.Font.Name), size)
		}
		if font == 0 {
			font = send(class("NSFont"), "systemFontOfSize:", size)
		}
		target := send(objc.ID(controllerClass), "new")
		activeFontPanel = &fontSelection{
			manager: manager, controller: target, font: send(font, "retain"),
			previousTarget: send(send(manager, "target"), "retain"),
			previousAction: objc.Send[objc.SEL](manager, selector("action")),
			onChange:       p.OnChange,
		}
		observeClose(target, panel, "pureKitFontPanelClosed:")
		send(manager, "setSelectedFont:isMultiple:", font, false)
		send(manager, "setTarget:", target)
		send(manager, "setAction:", selector("pureKitChangeFont:"))
		send(manager, "orderFrontFontPanel:", objc.ID(0))
		return nil
	})
}

func fontChanged(target objc.ID, _ objc.SEL, manager objc.ID) {
	s := activeFontPanel
	if s == nil || s.controller != target || manager != s.manager {
		return
	}
	font := send(manager, "convertFont:", s.font)
	if font == 0 {
		return
	}
	send(font, "retain")
	send(s.font, "release")
	s.font = font
	send(manager, "setSelectedFont:isMultiple:", font, false)
	if s.onChange != nil {
		s.onChange(Font{
			Name:   readString(send(font, "fontName")),
			Family: readString(send(font, "familyName")),
			Size:   objc.Send[float64](font, selector("pointSize")),
		})
	}
}

func fontPanelClosed(target objc.ID, _ objc.SEL, _ objc.ID) {
	if activeFontPanel != nil && activeFontPanel.controller == target {
		finishFontPanel()
	}
}

func finishFontPanel() {
	s := activeFontPanel
	if s == nil {
		return
	}
	activeFontPanel = nil
	send(notificationCenter(), "removeObserver:", s.controller)
	if send(s.manager, "target") == s.controller {
		send(s.manager, "setTarget:", s.previousTarget)
		send(s.manager, "setAction:", s.previousAction)
	}
	send(s.previousTarget, "release")
	send(s.font, "release")
	send(s.controller, "autorelease")
}
