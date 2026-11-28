package collectionview

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/display"
	"github.com/puremixlab/purekit/pkg/lucide"
)

func Basic(c *ui.Context) {
	root := ui.Column(c).Fill()
	state := ui.Local(root, "collection", func() display.CollectionState[int] {
		return display.CollectionState[int]{}
	})
	keys := make([]int, 48)
	for i := range keys {
		keys[i] = i + 1
	}
	name := func(item int) string { return fmt.Sprintf("Document %02d", item) }
	root.Children(func() {
		display.CollectionView(c, state, display.CollectionViewProps[int]{
			Label: "Project files", Items: keys, ItemLabel: name, Columns: 4, Multiple: true,
		}, func(item int) {
			ui.Column(c).Gap(4).AlignItems(ui.Center).Children(func() {
				ui.Icon(c, lucide.FileText()).Size(24, 24)
				ui.Text(c, name(item)).FontSize(12).SingleLine()
			})
		}).Grow(1)
	})
}
