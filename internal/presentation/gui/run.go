package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	screenAddSentenceSlug     = "screen1"
	screenProcessButchSlug    = "screen2"
	screenProcessOneByOneSlug = "main"
)

type App struct {
	app  *tview.Application
	root tview.Primitive

	pages *tview.Pages

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
	return a.app.
		SetRoot(a.root, true).
		SetFocus(a.input).
		Run()
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
	// основной экран (текущий)
	a.createAreas()
	a.createButtons()

	screenProcessOneByOne := a.createScreenProcessOneByOneGrid()
	screenAddSentence := a.createScreenAddSentenceGrid()
	screenProcessButch := a.createScreenProcessBatchGrid() // <-- новый экран вместо заглушки

	// pages со всеми экранами
	pages := tview.NewPages()
	pages.AddPage(screenAddSentenceSlug, screenAddSentence, true, false)
	pages.AddPage(screenProcessButchSlug, screenProcessButch, true, false)
	pages.AddPage(screenProcessOneByOneSlug, screenProcessOneByOne, true, true) // показываем главный по умолчанию
	a.pages = pages

	// тулбар сверху
	toolbar := a.createToolbar()

	// корневой layout: тулбар + pages
	root := tview.NewFlex()
	root.SetDirection(tview.FlexRow)
	root.AddItem(toolbar, 3, 1, false)
	root.AddItem(a.pages, 0, 1, true)

	a.root = root
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
		var err error
		text := a.input.GetText()
		if a.gen != nil {
			text, err = a.gen.GenerateNote(text, "")
			_ = err // TODO: обработка ошибки
		}
		a.dataView.SetText(text)
		a.input.SetText("", true)
		a.app.SetFocus(a.input)
	})

	a.saveButton = makeButton("Сохранить", func() {
		a.handleSave()
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

// общая логика сохранения — используется и на главном экране, и на Add sentence
func (a *App) handleSave() {
	text := a.input.GetText()
	_ = a.saver.Save(text) // TODO: обработка ошибки
	a.dataView.SetText("")
	a.input.SetText("", true)
	a.app.SetFocus(a.input)
}

// Экран 3: процессинг по одному (твой старый основной экран)
func (a *App) createScreenProcessOneByOneGrid() *tview.Grid {
	grid := tview.NewGrid()
	grid.SetRows(-1, -1, -1, -6)
	grid.SetColumns(-8, -2)
	grid.SetBorders(false)

	grid.AddItem(a.input, 0, 0, 3, 1, 3, 10, true)
	grid.AddItem(a.generateButton, 0, 1, 1, 1, 1, 10, false)
	grid.AddItem(a.saveButton, 1, 1, 1, 1, 1, 10, false)
	grid.AddItem(a.nextButton, 2, 1, 1, 1, 1, 10, false)
	grid.AddItem(a.dataView, 3, 0, 1, 2, 3, 10, false)

	return grid
}

// Экран 1: добавление предложения — большой input и кнопка Save в правом нижнем углу
func (a *App) createScreenAddSentenceGrid() *tview.Grid {
	grid := tview.NewGrid()
	grid.SetRows(-9, -1)
	grid.SetColumns(6)
	grid.SetBorders(false)

	saveBtn := tview.NewButton("Save")
	saveBtn.SetSelectedFunc(func() {
		a.handleSave()
	})
	saveBtn.SetBorder(true)
	saveBtn.SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack))
	saveBtn.SetLabelColor(tcell.ColorWhite)
	saveBtn.SetLabelColorActivated(tcell.ColorDarkGray)
	saveBtn.SetBackgroundColorActivated(tcell.ColorBlack)

	grid.AddItem(a.input, 0, 0, 1, 6, 0, 0, true)
	grid.AddItem(saveBtn, 1, 5, 1, 1, 0, 0, false)

	return grid
}

// Экран 2: batch — четыре колонки с кнопками и текстом, плюс кнопка ниже в последней колонке
func (a *App) createScreenProcessBatchGrid() *tview.Grid {
	grid := tview.NewGrid()
	grid.SetRows(-4, -1, -1, -3)
	// четыре колонки одинаковой ширины
	grid.SetColumns(-1, -1, -1, -1)
	grid.SetBorders(false)

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

	// Кнопки
	btn1 := makeButton("Copy first 10 sentences to clipboard", func() {})
	btn2 := makeButton("Add sentences to Anki from clipboard", func() {})
	btn3 := makeButton("Delete first 10 sentences", func() {})
	//btn4 := makeButton("Open file", func() {})

	// Текст во второй колонке (между кнопками)
	text := tview.NewTextView()
	text.SetTextAlign(tview.AlignCenter)
	text.SetText("Go to ChatGPT and generate CSV")

	// Первый «контентный» ряд (row 1)
	grid.AddItem(btn1, 1, 0, 1, 1, 0, 0, false) // можно дать фокус первой кнопке
	grid.AddItem(text, 1, 1, 1, 1, 0, 0, false)
	grid.AddItem(btn2, 1, 2, 1, 1, 0, 0, false)
	grid.AddItem(btn3, 1, 3, 1, 1, 0, 0, false)
	//grid.AddItem(btn4, 2, 3, 1, 1, 0, 0, false)

	return grid
}

// Плейсхолдер-экраны (если вдруг ещё пригодится)
func (a *App) createPlaceholderScreen(text string) tview.Primitive {
	tv := tview.NewTextView()
	tv.SetTextAlign(tview.AlignCenter)
	tv.SetText(text)

	grid := tview.NewGrid()
	grid.SetRows(0)
	grid.SetColumns(0)
	grid.SetBorders(false)
	grid.AddItem(tv, 0, 0, 1, 1, 0, 0, true)

	return grid
}

// Тулбар с тремя кнопками, переключающими страницы
func (a *App) createToolbar() *tview.Flex {
	styleButton := tcell.StyleDefault.Background(tcell.ColorBlack)

	makeTab := func(label string, handler func()) *tview.Button {
		btn := tview.NewButton(label)
		btn.SetSelectedFunc(handler)
		btn.SetBorder(true)
		btn.SetStyle(styleButton)
		btn.SetLabelColor(tcell.ColorWhite)
		btn.SetLabelColorActivated(tcell.ColorDarkGray)
		btn.SetBackgroundColorActivated(tcell.ColorBlack)
		return btn
	}

	btn1 := makeTab("Add sentence", func() {
		a.pages.SwitchToPage(screenAddSentenceSlug)
	})
	btn2 := makeTab("Process batch", func() {
		a.pages.SwitchToPage(screenProcessButchSlug)
	})
	btn3 := makeTab("Process one by one", func() {
		a.pages.SwitchToPage(screenProcessOneByOneSlug)
		a.app.SetFocus(a.input)
	})

	toolbar := tview.NewFlex()
	toolbar.SetDirection(tview.FlexColumn)
	toolbar.AddItem(btn1, 0, 1, false)
	toolbar.AddItem(btn2, 0, 1, false)
	toolbar.AddItem(btn3, 0, 1, false)
	toolbar.SetBorder(false)

	return toolbar
}
