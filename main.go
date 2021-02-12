package main

import (
	"github.com/awesome-gocui/gocui"
	"github.com/go-errors/errors"
	"log"
	"os"
)

func main() {
	os.Setenv("COLORTERM", "truecolor")
	g, err := gocui.NewGui(gocui.OutputTrue, false)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlD, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	g.BgColor = gocui.ColorBlack
	g.FgColor = gocui.ColorWhite
	drawHeader(g)
	drawBackground(g)
	drawText(g)
	drawInput(g)
	columnSize := 105 - 84
	drawDiagram(g, columnSize)
	drawInventory(g, columnSize)
	drawPreview(g, columnSize)
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
