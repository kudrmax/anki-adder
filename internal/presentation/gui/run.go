package gui

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

const (
	inputViewName  = "input"
	buttonViewName = "button"
)

type App struct {
	gui   *gocui.Gui
	saver Saver
}

func New(s Saver) (*App, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return nil, err
	}

	app := &App{
		gui:   g,
		saver: s,
	}

	g.Cursor = true
	g.Mouse = true
	g.SetManagerFunc(app.layout)

	err = app.SetKeybindings()
	if err != nil {
		return nil, err
	}

	return app, nil
}

// Run запускает главный цикл GUI.
func (a *App) Run() error {
	defer a.gui.Close()

	if err := a.gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}

	return nil
}

func (a *App) SetKeybindings() error {
	// Глобальный выход: Ctrl+C.
	if err := a.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		a.gui.Close()
		return err
	}

	// Submit: Ctrl+S в input.
	if err := a.gui.SetKeybinding(inputViewName, gocui.KeyCtrlS, gocui.ModNone, a.submitInput); err != nil {
		a.gui.Close()
		return err
	}

	// Submit: клик мышкой по кнопке.
	if err := a.gui.SetKeybinding(buttonViewName, gocui.MouseLeft, gocui.ModNone, a.submitInput); err != nil {
		a.gui.Close()
		return err
	}

	return nil
}

// layout — менеджер раскладки, рисует input и кнопку.
func (a *App) layout(g *gocui.Gui) error {

	err := a.layoutInput(a.gui)
	if err != nil {
		return err
	}

	err = a.layoutButton(a.gui)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) layoutInput(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	inputWidth := maxX - 2
	inputHeight := 10
	if inputWidth > maxX-2 {
		inputWidth = maxX - 2
	}
	if inputHeight > maxY-2 {
		inputHeight = maxY - 2
	}

	x0 := (maxX - inputWidth) / 2
	x1 := x0 + inputWidth
	y0 := 2
	y1 := y0 + inputHeight

	if v, err := g.SetView(inputViewName, x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " Input "
		v.Editable = true
		v.Editor = gocui.DefaultEditor

		if _, err := g.SetCurrentView(inputViewName); err != nil {
			return err
		}
	}

	return nil
}

// layout — менеджер раскладки, рисует input и кнопку.
func (a *App) layoutButton(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	btnWidth := 20
	if btnWidth > maxX-2 {
		btnWidth = maxX - 2
	}

	bx0 := (maxX - btnWidth) / 2
	bx1 := bx0 + btnWidth
	by0 := maxY/2 - 7/2 + 1 + 7
	by1 := by0 + 2

	if v, err := g.SetView(buttonViewName, bx0, by0, bx1, by1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		v.Editable = false

		fmt.Fprintln(v, "   Submit (Ctrl+S)   ")
	}

	return nil
}

// submitInput — вызывается при Ctrl+S в input или клике по кнопке.
func (a *App) submitInput(g *gocui.Gui, v *gocui.View) error {
	inputView, err := g.View(inputViewName)
	if err != nil {
		return err
	}

	// Берём весь буфер, включая переносы строк.
	text := strings.TrimSpace(inputView.Buffer())

	// Передаём текст внешнему Saver.
	a.saver.Save(text)

	// Очистить поле ввода и вернуть курсор в начало.
	inputView.Clear()
	if err := inputView.SetCursor(0, 0); err != nil {
		return err
	}

	return nil
}

func convertCoords(leftX, upY, lenX, lenY int) (x0, y0, x1, y1 int) {
	x0, y0 = leftX, upY
	x1 = x0 + lenX
	y1 = y0 + lenY
	return
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
