package desktop

import (
	"errors"
	"sync"

	"github.com/ebitengine/purego/objc"
	"github.com/egoist/mygo"
)

var shareDelegateClass = sync.OnceValues(func() (objc.Class, error) {
	return objc.RegisterClass("PureKitSharingPickerDelegate", objc.GetClass("NSObject"), nativeProtocols("NSSharingServicePickerDelegate"), nil, []objc.MethodDef{
		{Cmd: selector("sharingServicePicker:didChooseSharingService:"), Fn: shareServiceChosen},
		{Cmd: selector("pureKitShareWindowClosed:"), Fn: shareWindowClosed},
	})
})

// The picker and its weak delegate must stay alive until a service is chosen or dismissed.
var sharePicker, shareDelegate objc.ID

// ShowShare shows the macOS system share picker anchored to the parent content.
func ShowShare(parent *mygo.Window, p ShareProps) error {
	if p.Text == "" && len(p.Files) == 0 {
		return errors.New("provide text or files to share")
	}
	return mainThread(func() error {
		window, err := nativeWindow(parent)
		if err != nil {
			return err
		}
		view := send(window, "contentView")
		if view == 0 {
			return ErrWindow
		}
		if sharePicker != 0 {
			return errors.New("the share picker is already visible")
		}
		items, err := fileURLs(p.Files)
		if err != nil {
			return err
		}
		if p.Text != "" {
			send(items, "addObject:", nativeString(p.Text))
		}
		delegateClass, err := shareDelegateClass()
		if err != nil {
			return err
		}
		picker := send(send(class("NSSharingServicePicker"), "alloc"), "initWithItems:", items)
		if picker == 0 {
			return ErrUnavailable
		}
		sharePicker, shareDelegate = picker, send(objc.ID(delegateClass), "new")
		observeClose(shareDelegate, window, "pureKitShareWindowClosed:")
		send(picker, "setDelegate:", shareDelegate)
		bounds := objc.Send[nsRect](view, selector("bounds"))
		rect := nsRect{
			Origin: nsPoint{X: bounds.Origin.X + float64(p.Anchor.X), Y: bounds.Origin.Y + float64(p.Anchor.Y)},
			Size:   nsSize{Width: max(1, float64(p.Anchor.W)), Height: max(1, float64(p.Anchor.H))},
		}
		if !objc.Send[bool](view, selector("isFlipped")) {
			rect.Origin.Y = bounds.Origin.Y + bounds.Size.Height - float64(p.Anchor.Y) - rect.Size.Height
		}
		const minYEdge = uintptr(1)
		send(picker, "showRelativeToRect:ofView:preferredEdge:", rect, view, minYEdge)
		return nil
	})
}

func shareServiceChosen(delegate objc.ID, _ objc.SEL, picker, _ objc.ID) {
	if picker == sharePicker && delegate == shareDelegate {
		finishShare()
	}
}

func shareWindowClosed(delegate objc.ID, _ objc.SEL, _ objc.ID) {
	if delegate == shareDelegate {
		finishShare()
	}
}

func finishShare() {
	picker, delegate := sharePicker, shareDelegate
	sharePicker, shareDelegate = 0, 0
	send(notificationCenter(), "removeObserver:", delegate)
	send(picker, "setDelegate:", objc.ID(0))
	send(picker, "autorelease")
	send(delegate, "autorelease")
}
