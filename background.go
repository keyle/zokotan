package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
	"io/ioutil"
)

func drawBackground(g *gocui.Gui) {
	screenX, screenY := g.Size()
	if background, err := g.SetView("background", -1, 0, screenX, screenY, 0); err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}
		background.Frame = false
		//window.TitleColor = gocui.ColorMagenta
		//window.FrameColor = gocui.ColorMagenta
		background.FrameRunes = []rune{'x', '│'}
		background.Title = "View"
		background.Editable = false
		background.Wrap = false

		txt, err := ioutil.ReadFile("./ansi/jungle2.txt")
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(background, string(txt))

		//window.SetWritePos(3, 0)
		//fmt.Fprintln(window, "Welcome to Zokotan."+
		//	"\n\n  You have crash landed on an unknown planet whilst in transit to Cerca IV."+
		//	"\n\nMy name is IRis and I am your internal assistant."+
		//	"\n\nAs a brain chip, I will assist you through the rest of your life. \nI am here to help!"+
		//	"\n\nThere is a lush jungle outside, mostly purple. The sky is orange. \nYou can hear noise as there is a waterfall nearby, \nwhich appears to be liquid methane."+
		//	"\n\nThere is no sign of intelligent life on this planet, \nalthough there is a lot of animal activity."+
		//	"\n\nThe air appears to be breathable. You could take off your helmet.")
	}
}
