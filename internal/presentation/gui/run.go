package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type App struct {
	app  *tview.Application
	root *tview.Grid

	input    *tview.TextArea
	dataView *tview.TextView

	generateButton *tview.Button
	saveButton     *tview.Button
	nextButton     *tview.Button

	saver Saver
	gen   Generator
	next  NextProvider
}

func NewApp(s Saver, g Generator, n NextProvider) *App {
	a := &App{
		app:   tview.NewApplication().EnableMouse(true),
		saver: s,
		gen:   g,
		next:  n,
	}

	a.configureBorders()
	a.build()

	return a
}

func (a *App) Run() error {
	return a.app.SetRoot(a.root, true).SetFocus(a.input).Run()
}

func (a *App) configureBorders() {
	tview.Borders.HorizontalFocus = tview.Borders.Horizontal
	tview.Borders.VerticalFocus = tview.Borders.Vertical
	tview.Borders.TopLeftFocus = tview.Borders.TopLeft
	tview.Borders.TopRightFocus = tview.Borders.TopRight
	tview.Borders.BottomLeftFocus = tview.Borders.BottomLeft
	tview.Borders.BottomRightFocus = tview.Borders.BottomRight
}

func (a *App) build() {
	a.createAreas()
	a.createButtons()
	a.createGrid()
}

func (a *App) createAreas() {
	input := tview.NewTextArea()
	input.SetTitle(" Ввод текста ")
	input.SetTitleAlign(tview.AlignLeft)
	input.SetBorder(true)
	a.input = input

	dataView := tview.NewTextView()
	dataView.SetBorder(true)
	dataView.SetTitle(" Данные ")
	dataView.SetTitleAlign(tview.AlignLeft)
	a.dataView = dataView
}

func (a *App) createButtons() {
	styleButton := tcell.StyleDefault.Background(tcell.ColorBlack)

	makeButton := func(label string, handler func()) *tview.Button {
		btn := tview.NewButton(label)
		btn.SetSelectedFunc(handler)
		btn.SetBorder(true)
		btn.SetStyle(styleButton)
		btn.SetLabelColor(tcell.ColorWhite)
		btn.SetLabelColorActivated(tcell.ColorDarkGray)
		btn.SetBackgroundColorActivated(tcell.ColorBlack)
		return btn
	}

	a.generateButton = makeButton("Сгенерировать", func() {
		text := a.input.GetText()
		if a.gen != nil {
			text = a.gen.Generate(text)
		}
		a.dataView.SetText(text)
		a.input.SetText("", true)
		a.app.SetFocus(a.input)
	})

	a.saveButton = makeButton("Сохранить", func() {
		text := a.input.GetText()
		_ = a.saver.Save(text)
		a.dataView.SetText("")
		a.input.SetText("", true)
		a.app.SetFocus(a.input)
	})

	a.nextButton = makeButton("Следующее предложение", func() {
		var text string
		if a.next != nil {
			text = a.next.Next()
		}
		a.dataView.SetText("")
		a.input.SetText(text, true)
		a.app.SetFocus(a.input)
	})
}

func (a *App) createGrid() {
	grid := tview.NewGrid().
		SetRows(-1, -1, -1, -6).
		SetColumns(-8, -2).
		SetBorders(false)

	grid.AddItem(a.input, 0, 0, 3, 1, 3, 10, true)
	grid.AddItem(a.generateButton, 0, 1, 1, 1, 1, 10, false)
	grid.AddItem(a.saveButton, 1, 1, 1, 1, 1, 10, false)
	grid.AddItem(a.nextButton, 2, 1, 1, 1, 1, 10, false)
	grid.AddItem(a.dataView, 3, 0, 1, 2, 3, 10, false)

	a.root = grid
}
