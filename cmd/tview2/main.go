package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication().
		EnableMouse(true) // <-- ВКЛЮЧАЕМ МЫШЬ

	// --- Строка 2: контейнер "Данные" ---
	dataView := tview.NewTextView().
		SetTextAlign(tview.AlignLeft)
	dataView.SetBorder(true).
		SetTitle("Данные")

	// --- Строка 1: TextArea + кнопки ---
	input := tview.NewTextArea()
	input.SetBorder(true).
		SetTitle("Ввод текста")
	input.SetSize(3, 0) // высота 3 строки, ширина тянется Flex-ом

	// Кнопка 1: переносит текст в "Данные" и очищает ввод
	button1 := tview.NewButton("Кнопка 1")
	button1.SetSelectedFunc(func() {
		text := input.GetText()
		dataView.SetText(text)
		input.SetText("", false) // очистили поле
		app.SetFocus(input)      // вернули фокус в поле ввода (по желанию)
	})

	// Кнопка 2: пока заглушка
	button2 := tview.NewButton("Кнопка 2")
	button2.SetSelectedFunc(func() {
		dataView.SetText("Нажата Кнопка 2 (пока заглушка)")
	})

	// Горизонтальный Flex для первой строки
	row1 := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(input, 0, 3, true). // input растягивается
		AddItem(button1, 12, 0, false).
		AddItem(button2, 12, 0, false)

	// Вертикальный Flex: строка 1 сверху, "Данные" снизу
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(row1, 5, 0, true).     // 5 строк высота под первую строку
		AddItem(dataView, 0, 1, false) // остальное под "Данные"

	if err := app.SetRoot(layout, true).SetFocus(input).Run(); err != nil {
		panic(err)
	}
}
