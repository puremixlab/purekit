package desktop

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"structs"
	"sync"

	"github.com/ebitengine/purego/objc"
	"github.com/egoist/mygo"
)

type nsPoint struct {
	_    structs.HostLayout
	X, Y float64
}

type nsSize struct {
	_             structs.HostLayout
	Width, Height float64
}

type nsRect struct {
	_      structs.HostLayout
	Origin nsPoint
	Size   nsSize
}

var selectors sync.Map

func selector(name string) objc.SEL {
	if existing, ok := selectors.Load(name); ok {
		return existing.(objc.SEL)
	}
	sel := objc.RegisterName(name)
	selectors.Store(name, sel)
	return sel
}

func class(name string) objc.ID { return objc.ID(objc.GetClass(name)) }

func send(receiver objc.ID, method string, args ...any) objc.ID {
	return receiver.Send(selector(method), args...)
}

func nativeString(value string) objc.ID {
	return send(class("NSString"), "stringWithUTF8String:", value)
}

func readString(value objc.ID) string {
	if value == 0 {
		return ""
	}
	return objc.Send[string](value, selector("UTF8String"))
}

func mainThread(work func() error) error {
	if err := nativeReady(); err != nil {
		return err
	}
	var result error = ErrNotReady
	mygo.RunOnMain(func() { result = work() })
	return result
}

func nativeWindow(parent *mygo.Window) (objc.ID, error) {
	if parent == nil {
		return 0, ErrWindow
	}
	window := objc.ID(parent.NativeHandle())
	if window == 0 {
		return 0, ErrWindow
	}
	return window, nil
}

func fileURLs(paths []string) (objc.ID, error) {
	items := send(class("NSMutableArray"), "array")
	for _, path := range paths {
		if path == "" || strings.ContainsRune(path, 0) {
			return 0, errors.New("a file path is empty or invalid")
		}
		full, err := filepath.Abs(path)
		if err != nil {
			return 0, err
		}
		url := send(class("NSURL"), "fileURLWithPath:", nativeString(full))
		if url == 0 {
			return 0, fmt.Errorf("could not create a file URL for %q", path)
		}
		send(items, "addObject:", url)
	}
	return items, nil
}

func notificationCenter() objc.ID {
	return send(class("NSNotificationCenter"), "defaultCenter")
}

func observeClose(observer, window objc.ID, method string) {
	send(notificationCenter(), "addObserver:selector:name:object:",
		observer, selector(method), nativeString("NSWindowWillCloseNotification"), window)
}
