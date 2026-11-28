package segmented

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/action"
)

var period int

func Basic(c *ui.Context) {
	action.Segmented(c, &period, []string{"Daily", "Weekly", "Monthly", "Yearly"}, action.SegmentedProps{})
}
