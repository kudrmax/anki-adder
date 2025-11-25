package main

import (
	"fmt"
	"math/rand"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication().EnableMouse(true)

	tview.Borders.HorizontalFocus = tview.Borders.Horizontal
	tview.Borders.VerticalFocus = tview.Borders.Vertical
	tview.Borders.TopLeftFocus = tview.Borders.TopLeft
	tview.Borders.TopRightFocus = tview.Borders.TopRight
	tview.Borders.BottomLeftFocus = tview.Borders.BottomLeft
	tview.Borders.BottomRightFocus = tview.Borders.BottomRight

	input := tview.NewTextArea()
	input.SetTitle(" Ввод текста ")
	input.SetTitleAlign(tview.AlignLeft)
	input.SetBorder(true)

	dataView := tview.NewTextView()
	dataView.SetBorder(true)
	dataView.SetTitle(" Данные ")
	dataView.SetTitleAlign(tview.AlignLeft)

	generateButton := tview.NewButton("Сгенерировать")
	generateButton.SetSelectedFunc(func() {
		text := input.GetText()
		text = ProcessGenerate(text)

		dataView.SetText(text)
		input.SetText("", true)

		app.SetFocus(input)
	})
	generateButton.SetBorder(true)
	generateButton.SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack))
	generateButton.SetLabelColor(tcell.ColorWhite)
	generateButton.SetLabelColorActivated(tcell.ColorDarkGray)
	generateButton.SetBackgroundColorActivated(tcell.ColorBlack)

	saveButton := tview.NewButton("Сохранить")
	saveButton.SetSelectedFunc(func() {
		text := input.GetText()
		ProcessSave(text)

		dataView.SetText("")
		input.SetText("", true)
		app.SetFocus(input)
	})
	saveButton.SetBorder(true)
	saveButton.SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack))
	saveButton.SetLabelColor(tcell.ColorWhite)
	saveButton.SetLabelColorActivated(tcell.ColorDarkGray)
	saveButton.SetBackgroundColorActivated(tcell.ColorBlack)

	nextButton := tview.NewButton("Следующее предложение")
	nextButton.SetSelectedFunc(func() {
		text := ProcessNext()

		dataView.SetText("")
		input.SetText(text, true)
		app.SetFocus(input)
	})
	nextButton.SetBorder(true)
	nextButton.SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack))
	nextButton.SetLabelColor(tcell.ColorWhite)
	nextButton.SetLabelColorActivated(tcell.ColorDarkGray)
	nextButton.SetBackgroundColorActivated(tcell.ColorBlack)

	grid := tview.NewGrid().
		SetRows(-1, -1, -1, -6).
		SetColumns(-8, -2).
		SetBorders(false)

	grid.AddItem(input, 0, 0, 3, 1, 3, 10, true)
	grid.AddItem(generateButton, 0, 1, 1, 1, 1, 10, false)
	grid.AddItem(saveButton, 1, 1, 1, 1, 1, 10, false)
	grid.AddItem(nextButton, 2, 1, 1, 1, 1, 10, false)
	grid.AddItem(dataView, 3, 0, 1, 2, 3, 10, false)

	if err := app.SetRoot(grid, true).SetFocus(input).Run(); err != nil {
		panic(err)
	}
}

func ProcessGenerate(text string) string {
	// TODO: пока заглушка
	return fmt.Sprintf("Сгенерировано: %s", text)
}

func ProcessSave(text string) {
	// TODO: пока заглушка
}

func ProcessNext() string {
	// TODO: пока заглушка
	return fmt.Sprintf("Новое предложение номер: %d", rand.Int())
}
