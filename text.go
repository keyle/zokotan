package main

import (
	"fmt"
	"github.com/awesome-gocui/gocui"
)

func drawText(g *gocui.Gui) {
	screenX, _ := g.Size()
	txtWidth := 83
	if text, err := g.SetView("text", screenX/6*1, 1, screenX/6*1+txtWidth, 35, 0); err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}

		if _, err := g.SetCurrentView("text"); err != nil {
			panic(err)
		}
		text.Frame = false
		text.Autoscroll = true
		//text.Editable = false
		//text.Wrap = true
		//text.Highlight = false
		//text.Autoscroll = false

		fmt.Fprintln(text, "Welcome to Zokotan."+
			"\n\nYou have crash landed on an unknown planet whilst in transit to Cerca IV."+
			"\n\nMy name is IRis and I am your internal assistant."+
			"\n\nAs a brain chip, I will assist you through the rest of your life. \nI am here to help!"+
			"\n\nThere is a lush jungle outside, mostly purple. The sky is orange. \nYou can hear noise as there is a waterfall nearby, \nwhich appears to be liquid methane."+
			"\n\nThere is no sign of intelligent life on this planet, \nalthough there is a lot of animal activity."+
			"\n\nThe air appears to be breathable. You could take off your helmet.")
	}
}
