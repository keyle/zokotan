package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
)

func drawInventory(g *gocui.Gui, columnSize int) {
	screenX, _ := g.Size()
	if inventory, err := g.SetView("inventory", screenX/6*5, 9, screenX/6*5+columnSize, 17, 0); err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}
		inventory.Frame = true
		inventory.TitleColor = gocui.ColorCyan
		inventory.FrameColor = gocui.ColorCyan
		inventory.Title = "Inventory (4/5)"
		inventory.Editable = false
		inventory.Wrap = false

		fmt.Fprintln(inventory, "               ")
		fmt.Fprintln(inventory, " Helmet        ")
		fmt.Fprintln(inventory, " Suit          ")
		fmt.Fprintln(inventory, " Gloves        ")
		fmt.Fprintln(inventory, " Travicom      ")
	}
}
