package display

import (
	"fmt"
	"slices"
	"testing"

	"github.com/egoist/mygo/ui"
)

func TestCollectionNavigationScrollsToOffscreenTiles(t *testing.T) {
	items := make([]int, 200)
	for i := range items {
		items[i] = i
	}
	var state CollectionState[int]
	name := func(i int) string { return fmt.Sprintf("Item %d", i) }
	tester := ui.NewTester(func(c *ui.Context) {
		CollectionView(c, &state, CollectionViewProps[int]{
			Items: items, ItemLabel: name, Columns: 3, ItemHeight: 54,
		}, func(i int) { ui.Text(c, name(i)) })
	}, 360, 210)
	if tester.HasText(name(199)) {
		t.Fatal("the last tile was built in the initial viewport")
	}
	if err := tester.Click(name(0)); err != nil {
		t.Fatal(err)
	}
	tester.Key(0, ui.KeyRight)
	if !state.Selection.Has(1) || state.Selection.Has(0) {
		t.Fatal("Right did not move selection to the next tile")
	}
	tester.Key(0, ui.KeyEnd)
	tester.Frame()
	if !state.Selection.Has(199) || !tester.HasText(name(199)) {
		t.Fatalf("End did not reach the last tile: %v", tester.Texts())
	}
}

func TestColumnViewChangesThePathWhenChoosingAnotherBranch(t *testing.T) {
	state := ColumnViewState[string]{Path: []string{"A", "One"}}
	children := map[string][]string{"A": {"One"}, "B": {"Two"}}
	tester := ui.NewTester(func(c *ui.Context) {
		ColumnView(c, &state, ColumnViewProps[string]{
			Roots: []string{"A", "B"}, Children: func(key string) []string { return children[key] },
		}, func(key string) { ui.Text(c, key) })
	}, 440, 220)
	if err := tester.Click("B"); err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(state.Path, []string{"B"}) || !tester.HasText("Two") || tester.HasText("One") {
		t.Fatalf("path %v, texts %v", state.Path, tester.Texts())
	}
}
