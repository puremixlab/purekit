package stepper

import (
	"fmt"

	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/field"
)

func Basic(c *ui.Context) {
	root := ui.Row(c).AlignItems(ui.Center).Gap(12)
	copies := ui.Local(root, "copies", func() float64 { return 1 })
	root.Children(func() {
		ui.Text(c, fmt.Sprintf("%.0f copies", *copies)).Width(78)
		field.Stepper(c, copies, field.StepperProps{Label: "Copies", Min: 1, Max: 99})
	})
}
