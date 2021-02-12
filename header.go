package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
)

func drawHeader(g *gocui.Gui) {
	screenX, _ := g.Size()
	if textHeader, err := g.SetView("header", -1, -1, screenX, 1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}
		textHeader.BgColor = gocui.Get256Color(242)
		textHeader.FgColor = gocui.ColorBlack
		textHeader.Frame = false
		textHeader.Editable = false
		textHeader.Wrap = false

		fmt.Fprintln(textHeader, "    Crash Site")
	}
}
