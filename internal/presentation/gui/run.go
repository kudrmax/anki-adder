package gui

import (
	"fmt"
	"log"
	"strings"

	"github.com/jroimartin/gocui"
)

type GUI struct {
	g *gocui.Gui
}

func New() *GUI {
	gui := &GUI{}

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = true

	g.SetManagerFunc(gui.layout)

	// Global quit keybinding: Ctrl+C
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.quit); err != nil {
		log.Panicln(err)
	}

	// Submit with Enter when focus is on the input line
	if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, gui.submitInput); err != nil {
		log.Panicln(err)
	}

	// Submit with mouse click on the button
	if err := g.SetKeybinding("button", gocui.MouseLeft, gocui.ModNone, gui.submitInput); err != nil {
		log.Panicln(err)
	}

	gui.g = g
	return gui
}

func (gui *GUI) Run() {
	if err := gui.g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func (gui *GUI) layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// ---- Input view (centered) ----
	inputWidth := maxX / 2
	inputHeight := 2

	x0 := (maxX - inputWidth) / 2
	x1 := x0 + inputWidth
	y0 := (maxY - inputHeight) / 2
	y1 := y0 + inputHeight

	if v, err := g.SetView("input", x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " Write a sentence "
		v.Editable = true
		v.Editor = gocui.DefaultEditor
		v.Frame = true

		if _, err := g.SetCurrentView("input"); err != nil {
			return err
		}
	}

	// ---- Button view (under the input) ----
	btnWidth := 13
	if btnWidth > maxX-2 {
		btnWidth = maxX - 2
	}

	bx0 := (maxX - btnWidth) / 2
	bx1 := bx0 + btnWidth
	by0 := y1 + 1
	by1 := by0 + 2

	if v, err := g.SetView("button", bx0, by0, bx1, by1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		v.Editable = false

		// Text inside the button view
		fmt.Fprintln(v, "   Submit   ")
	}

	return nil
}

// submitInput is called when the user presses Enter in the input
// or clicks on the button with the mouse.
func (gui *GUI) submitInput(g *gocui.Gui, v *gocui.View) error {
	inputView, err := g.View("input")
	if err != nil {
		return err
	}

	// Read and trim the input text
	text := strings.TrimSpace(inputView.Buffer())

	// Call your stub function
	gui.doSomething(text)

	// Clear input and reset cursor
	inputView.Clear()
	if err := inputView.SetCursor(0, 0); err != nil {
		return err
	}

	return nil
}

// doSomething is a stub where you can put your own logic.
func (gui *GUI) doSomething(s string) {
	//log.Printf("doSomething called with: %q", s)
}

func (gui *GUI) quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
