package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
)

func drawDiagram(g *gocui.Gui, columnSize int) {
	screenX, _ := g.Size()
	if diagram, err := g.SetView("diagram", screenX/6*5, 2, screenX/6*5+columnSize, 8, 0); err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}
		diagram.Frame = true
		diagram.FgColor = gocui.ColorWhite
		diagram.TitleColor = gocui.ColorGreen
		diagram.FrameColor = gocui.ColorGreen
		diagram.Title = "Directions"
		diagram.Editable = false
		diagram.Wrap = false

		fmt.Fprintln(diagram, "              ")
		fmt.Fprintln(diagram, "         ⬆    ")
		fmt.Fprintln(diagram, "      ⬅     ⇨")
		fmt.Fprintln(diagram, "         ⬇    ")
	}
}
