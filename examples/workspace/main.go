package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/catalog/workspace"
	"github.com/puremixlab/purekit/pkg/layout"
	"github.com/puremixlab/purekit/pkg/theme"
)

func main() {
	state := workspace.New()
	configDir, err := os.UserConfigDir()
	var path string
	if err == nil {
		path = filepath.Join(configDir, "PureKit Workspace", "layout.json")
		readLayout(path, &state.Layout)
	}
	mygo.App.WhenReady(func() {
		mygo.App.SetMenu(state.NativeMenu())
		window := mygo.NewWindow(mygo.WindowOptions{
			Title: "PureKit Workspace", Width: 1000, Height: 650,
			MinWidth: 580, MinHeight: 360,
			Content: ui.View(func(c *ui.Context) {
				theme.Apply(c)
				state.View(c)
			}),
		})
		if path != "" {
			window.OnClosed(func() {
				if err := saveLayout(path, state.Layout); err != nil {
					log.Printf("save layout: %v", err)
				}
			})
		}
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}

func readLayout(path string, layoutState *layout.SplitViewState) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	var saved layout.SplitViewState
	if err == nil {
		err = json.Unmarshal(data, &saved)
	}
	if err != nil {
		log.Printf("read layout: %v", err)
		return
	}
	*layoutState = saved
}

func saveLayout(path string, layoutState layout.SplitViewState) error {
	data, err := json.Marshal(layoutState)
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	file, err := os.CreateTemp(dir, "layout-*.json")
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())
	_, writeErr := file.Write(data)
	if err := errors.Join(writeErr, file.Close()); err != nil {
		return err
	}
	return os.Rename(file.Name(), path)
}
