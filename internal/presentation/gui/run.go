package gui

import (
	"fmt"
	"log"
	"strings"
	"unicode"

	"github.com/jroimartin/gocui"
)

func Run() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = true

	g.SetManagerFunc(layout)

	// Global quit keybinding: Ctrl+C
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	// Submit: Ctrl+S в input
	if err := g.SetKeybinding("input", gocui.KeyCtrlS, gocui.ModNone, submitInput); err != nil {
		log.Panicln(err)
	}

	// Submit: клик мышкой по кнопке
	if err := g.SetKeybinding("button", gocui.MouseLeft, gocui.ModNone, submitInput); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// ---- Многострочный input по центру ----
	inputWidth := 50
	inputHeight := 7 // можно видеть несколько строк
	if inputWidth > maxX-2 {
		inputWidth = maxX - 2
	}

	x0 := (maxX - inputWidth) / 2
	x1 := x0 + inputWidth
	y0 := maxY/2 - inputHeight/2
	y1 := y0 + inputHeight

	if v, err := g.SetView("input", x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = " Input (Enter = newline, Ctrl+S = submit) "
		v.Editable = true
		v.Editor = gocui.EditorFunc(customEditor)

		if _, err := g.SetCurrentView("input"); err != nil {
			return err
		}
	}

	// ---- Кнопка под инпутом ----
	btnWidth := 20
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
		v.Title = " [ Button ] "
		v.Editable = false

		fmt.Fprintln(v, "   Submit (Ctrl+S)   ")
	}

	return nil
}

// customEditor — удобное редактирование текста (стрелки, Home/End, Ctrl+U/K/W и т.п.)
func customEditor(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case ch != 0 && mod == 0:
		// обычные символы
		v.EditWrite(ch)

	case key == gocui.KeySpace:
		v.EditWrite(' ')

	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)

	case key == gocui.KeyDelete:
		v.EditDelete(false)

	case key == gocui.KeyArrowLeft:
		v.MoveCursor(-1, 0, false)

	case key == gocui.KeyArrowRight:
		v.MoveCursor(1, 0, false)

	// Home / Ctrl+A — в начало строки
	case key == gocui.KeyHome || key == gocui.KeyCtrlA:
		cx, _ := v.Cursor()
		if cx > 0 {
			v.MoveCursor(-cx, 0, false)
		}

	// End / Ctrl+E — в конец строки
	case key == gocui.KeyEnd || key == gocui.KeyCtrlE:
		_, cy := v.Cursor()
		line, err := v.Line(cy)
		if err != nil {
			return
		}
		line = strings.TrimRight(line, "\n")
		runes := []rune(line)
		cx, _ := v.Cursor()
		diff := len(runes) - cx
		if diff > 0 {
			v.MoveCursor(diff, 0, false)
		}

	// Ctrl+U — очистить всю текущую строку
	case key == gocui.KeyCtrlU:
		_, cy := v.Cursor()
		line, err := v.Line(cy)
		if err != nil {
			return
		}
		line = strings.TrimRight(line, "\n")
		runes := []rune(line)
		for range runes {
			v.EditDelete(true)
		}

	// Ctrl+K — удалить от курсора до конца строки
	case key == gocui.KeyCtrlK:
		_, cy := v.Cursor()
		line, err := v.Line(cy)
		if err != nil {
			return
		}
		line = strings.TrimRight(line, "\n")
		runes := []rune(line)
		cx, _ := v.Cursor()
		toDelete := len(runes) - cx
		for i := 0; i < toDelete; i++ {
			v.EditDelete(false)
		}

	// Ctrl+W — удалить предыдущее слово
	case key == gocui.KeyCtrlW:
		cx, cy := v.Cursor()
		if cx == 0 {
			return
		}
		line, err := v.Line(cy)
		if err != nil {
			return
		}
		line = strings.TrimRight(line, "\n")
		runes := []rune(line)
		i := cx - 1

		// сначала удаляем пробелы слева
		for i >= 0 && unicode.IsSpace(runes[i]) {
			v.EditDelete(true)
			i--
		}
		// затем само слово
		for i >= 0 && !unicode.IsSpace(runes[i]) {
			v.EditDelete(true)
			i--
		}

	// Enter — теперь добавляет новую строку (для многострочного текста)
	case key == gocui.KeyEnter:
		v.EditNewLine()

	default:
		// остальные клавиши игнорируем
	}
}

// submitInput — вызывается при Ctrl+S в input или клике по кнопке.
func submitInput(g *gocui.Gui, v *gocui.View) error {
	inputView, err := g.View("input")
	if err != nil {
		return err
	}

	// Берём весь буфер, включая переносы строк
	text := strings.TrimSpace(inputView.Buffer())

	doSomething(text)

	// Очистить поле ввода и вернуть курсор в начало
	inputView.Clear()
	if err := inputView.SetCursor(0, 0); err != nil {
		return err
	}

	return nil
}

// doSomething — заглушка, сюда потом вставишь свою логику.
func doSomething(s string) {
	log.Printf("doSomething called with:\n%s\n---", s)
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
