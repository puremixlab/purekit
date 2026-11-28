package tokenfield

import (
	"github.com/egoist/mygo/ui"
	"github.com/puremixlab/purekit/pkg/field"
)

func Basic(c *ui.Context) {
	root := ui.Column(c).Width(320).Gap(8)
	tags := ui.Local(root, "tags", func() []string { return []string{"Desktop", "Design"} })
	root.Children(func() {
		ui.Text(c, "Labels").FontWeight(600)
		field.TokenField(c, tags, field.TokenFieldProps{
			Label: "Labels", Suggestions: []string{"Desktop", "Design", "Research", "Engineering", "Release"},
		})
	})
}
