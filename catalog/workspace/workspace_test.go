package workspace

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/display"
)

func TestWorkspaceSelectionAndCommands(t *testing.T) {
	state := New()
	menu := state.NativeMenu()
	preview := menu.ItemByID("preview")
	tester := ui.NewTester(state.View, 720, 360)
	if preview.IsEnabled() {
		t.Fatal("preview is enabled without a selection")
	}
	clickText(t, tester, "Design")
	if !tester.HasText("Brand guide.pdf") {
		t.Fatalf("folder selection did not change the table: %v", tester.Texts())
	}
	clickText(t, tester, "Brand guide.pdf")
	if !preview.IsEnabled() {
		t.Fatal("native preview did not follow the selected row")
	}
	tester.Key(0, ui.KeyEnter)
	tester.Advance(time.Second)
	if !tester.HasText("Document preview") {
		t.Fatal("Enter on the selected row did not open the preview")
	}
}

func TestWorkspaceShortcutRestoresInspectorWidth(t *testing.T) {
	state := New()
	state.Layout.DetailPercent = 60
	tester := ui.NewTester(state.View, 720, 360)
	mods, key, ok := display.HotkeyBinding(state.commands[1].Shortcut, display.HotkeyAuto)
	if !ok {
		t.Fatal("inspector command has no shortcut")
	}
	tester.Key(mods, key)
	if !state.Layout.InspectorHidden || tester.HasText("Details") {
		t.Fatal("shortcut did not hide the inspector")
	}
	tester.Key(mods, key)
	if state.Layout.InspectorHidden || !tester.HasText("Details") || state.Layout.DetailPercent != 60 {
		t.Fatal("shortcut did not restore the inspector at its previous size")
	}
}

func clickText(t *testing.T, tester *ui.Tester, name string) {
	t.Helper()
	box, ok := tester.Find(name)
	if !ok {
		t.Fatalf("could not find %q in %v", name, tester.Texts())
	}
	x, y := box.X+box.W/2, box.Y+box.H/2
	tester.Press(x, y)
	tester.Release(x, y)
}
