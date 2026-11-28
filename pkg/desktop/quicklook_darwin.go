package desktop

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
	"github.com/egoist/mygo"
)

type quickLookSession struct {
	window, panel, controller, previousResponder, urls objc.ID
	count                                              int
}

var activeQuickLook *quickLookSession

var loadQuickLook = sync.OnceValue(func() error {
	if class("QLPreviewPanel") != 0 {
		return nil
	}
	var loadErr error
	for _, library := range []string{
		"/System/Library/Frameworks/Quartz.framework/Frameworks/QuickLookUI.framework/QuickLookUI",
		"/System/Library/Frameworks/QuickLookUI.framework/QuickLookUI",
	} {
		if _, err := purego.Dlopen(library, purego.RTLD_LAZY|purego.RTLD_LOCAL); err != nil {
			loadErr = errors.Join(loadErr, err)
		} else if class("QLPreviewPanel") != 0 {
			return nil
		}
	}
	return fmt.Errorf("%w: Quick Look: %v", ErrUnavailable, loadErr)
})

var quickLookControllerClass = sync.OnceValues(func() (objc.Class, error) {
	return objc.RegisterClass("PureKitQuickLookController", objc.GetClass("NSResponder"), nativeProtocols("QLPreviewPanelDataSource"), nil, []objc.MethodDef{
		{Cmd: selector("acceptsPreviewPanelControl:"), Fn: acceptsPreview},
		{Cmd: selector("beginPreviewPanelControl:"), Fn: beginPreview},
		{Cmd: selector("endPreviewPanelControl:"), Fn: endPreview},
		{Cmd: selector("numberOfPreviewItemsInPreviewPanel:"), Fn: previewCount},
		{Cmd: selector("previewPanel:previewItemAtIndex:"), Fn: previewItem},
		{Cmd: selector("pureKitPreviewWindowClosed:"), Fn: previewWindowClosed},
	})
})

// ShowQuickLook previews one or more local files in the native Quick Look panel.
func ShowQuickLook(parent *mygo.Window, paths []string) error {
	if len(paths) == 0 {
		return errors.New("provide at least one file to preview")
	}
	return mainThread(func() error {
		window, err := nativeWindow(parent)
		if err != nil {
			return err
		}
		if err := loadQuickLook(); err != nil {
			return err
		}
		urls, err := fileURLs(paths)
		if err != nil {
			return err
		}
		controllerClass, err := quickLookControllerClass()
		if err != nil {
			return err
		}
		panel := send(class("QLPreviewPanel"), "sharedPreviewPanel")
		if panel == 0 {
			return ErrUnavailable
		}
		if s := activeQuickLook; s != nil && s.window == window {
			send(urls, "retain")
			send(s.urls, "release")
			s.urls, s.count = urls, len(paths)
		} else {
			finishQuickLook(true)
			controller := send(objc.ID(controllerClass), "new")
			previous := send(send(window, "nextResponder"), "retain")
			activeQuickLook = &quickLookSession{
				window: send(window, "retain"), panel: panel, controller: controller, previousResponder: previous,
				urls: send(urls, "retain"), count: len(paths),
			}
			send(controller, "setNextResponder:", previous)
			send(window, "setNextResponder:", controller)
			observeClose(controller, panel, "pureKitPreviewWindowClosed:")
			observeClose(controller, window, "pureKitPreviewWindowClosed:")
		}
		send(window, "makeKeyWindow")
		send(panel, "updateController")
		alreadyVisible := objc.Send[bool](panel, selector("isVisible"))
		send(panel, "makeKeyAndOrderFront:", objc.ID(0))
		if s := activeQuickLook; s == nil || send(panel, "currentController") != s.controller {
			finishQuickLook(!alreadyVisible)
			return errors.New("the parent window could not take control of Quick Look")
		}
		send(panel, "reloadData")
		send(panel, "setCurrentPreviewItemIndex:", int64(0))
		return nil
	})
}

func acceptsPreview(self objc.ID, _ objc.SEL, panel objc.ID) bool {
	s := activeQuickLook
	return s != nil && s.controller == self && s.panel == panel
}

func beginPreview(self objc.ID, _ objc.SEL, panel objc.ID) {
	if acceptsPreview(self, 0, panel) {
		send(panel, "setDataSource:", self)
	}
}

func endPreview(self objc.ID, _ objc.SEL, panel objc.ID) {
	if send(panel, "dataSource") == self {
		send(panel, "setDataSource:", objc.ID(0))
	}
}

func previewCount(self objc.ID, _ objc.SEL, panel objc.ID) int64 {
	if acceptsPreview(self, 0, panel) {
		return int64(activeQuickLook.count)
	}
	return 0
}

func previewItem(self objc.ID, _ objc.SEL, panel objc.ID, index int64) objc.ID {
	if acceptsPreview(self, 0, panel) && index >= 0 && index < int64(activeQuickLook.count) {
		return send(activeQuickLook.urls, "objectAtIndex:", uintptr(index))
	}
	return 0
}

func previewWindowClosed(self objc.ID, _ objc.SEL, note objc.ID) {
	if s := activeQuickLook; s != nil && s.controller == self {
		finishQuickLook(send(note, "object") == s.window)
	}
}

func finishQuickLook(closePanel bool) {
	s := activeQuickLook
	if s == nil {
		return
	}
	activeQuickLook = nil
	send(notificationCenter(), "removeObserver:", s.controller)
	endPreview(s.controller, 0, s.panel)
	if send(s.window, "nextResponder") == s.controller {
		send(s.window, "setNextResponder:", s.previousResponder)
	}
	if closePanel {
		send(s.panel, "close")
	}
	send(s.panel, "updateController")
	send(s.previousResponder, "release")
	send(s.window, "release")
	send(s.urls, "release")
	send(s.controller, "autorelease")
}
