package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"log"
)

func drawInput(g *gocui.Gui) {
	screenX, screenY := g.Size()
	if input, err := g.SetView("input", -1, screenY-2, screenX, screenY-0, 0); err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}
		input.BgColor = gocui.Get256Color(54)
		input.FgColor = gocui.ColorWhite
		input.Frame = false
		input.Editable = true
		input.Wrap = false
		g.Cursor = true
		input.Editor = gocui.EditorFunc(singleLineEditor)
		if err := g.SetKeybinding(input.Name(), gocui.KeyEnter, gocui.ModNone, parseMessage); err != nil {
			log.Panicln(err)
		}

		_, err := g.SetCurrentView(input.Name())
		if err != nil {
			panic(err)
		}
	}
}

func parseMessage(g *gocui.Gui, v *gocui.View) error {
	_, y := v.Cursor()
	text, _ := v.Line(y)
	v.Clear()
	v.Rewind()
	v.SetCursor(0, y)
	v.MoveCursor(0, 1)

	fmt.Println("\tYou have asked:", text)
	return nil
}
