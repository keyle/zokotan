package main

import (
	"github.com/awesome-gocui/gocui" // note this is fork
)

// singleLineEditor is used as a one line editor based off gocui.simpleEditor
func singleLineEditor(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case ch != 0 && mod == 0:
		v.EditWrite(ch)
	case key == gocui.KeySpace:
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
	case key == gocui.KeyDelete:
		v.EditDelete(false)
	case key == gocui.KeyEnter:
		return
	case key == gocui.KeyArrowDown:
		return
	case key == gocui.KeyArrowUp:
		return
	case key == gocui.KeyArrowLeft:
		v.MoveCursor(-1, 0)
	case key == gocui.KeyArrowRight:
		v.MoveCursor(1, 0)
	case key == gocui.KeyTab:
		return
	case key == gocui.KeySpace:
		v.EditWrite(' ')
	case key == gocui.KeyInsert:
		return
	default:
		v.EditWrite(ch)
	}
}
