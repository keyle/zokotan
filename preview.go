package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"io/ioutil"
)

func drawPreview(g *gocui.Gui, columnSize int) {
	screenX, _ := g.Size()
	if preview, err := g.SetView("preview", screenX/6*5, 18, screenX/6*5+columnSize, 9+18, 0); err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}
		preview.Frame = true
		preview.TitleColor = gocui.ColorYellow
		preview.FrameColor = gocui.ColorYellow
		preview.Title = "Map"
		preview.Editable = false
		preview.Wrap = false

		// https://manytools.org/hacker-tools/convert-images-to-ascii-art/
		// 200 wide, use color, invert true, download ansi.
		txt, err := ioutil.ReadFile("./ansi/jungle.txt")
		if err != nil {
			panic(err)
		}

		fmt.Fprintln(preview, string(txt))
	}

}
