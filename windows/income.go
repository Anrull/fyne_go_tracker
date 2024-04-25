package windows

import (
	"awesomeProject2/db"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

func WindowIncome() {
	profile := widget.NewLabel("Доходы")
	profile.Alignment = fyne.TextAlignCenter

	incomes := db.GetProfit("")

	buttonAdd := widget.NewButton("Добавить", func() { WindowAddIncome() })

	var options = []string{"Все позиции"}
	for _, p := range incomes {
		flag := true
		for _, elem := range options {
			if elem == p.Category {
				flag = false
				break
			}
		}
		if flag {
			options = append(options, p.Category)
		}
	}

	if len(incomes) > 0 {
		dropdown := widget.NewSelect(options, func(elem string) { status = elem })
		dropdown.SetSelectedIndex(0)

		list := widget.NewLabel("")

		wb := &strings.Builder{}
		for _, p := range incomes {
			wb.WriteString(fmt.Sprintf("%s %d\n", p.Category, p.Amount))
		}
		list.SetText(wb.String())

		dropdown.OnChanged = func(elem string) {
			if elem == "Все позиции" {
				elem = ""
			}
			outlays := db.GetProfit(elem)
			wb := &strings.Builder{}
			for _, p := range outlays {
				wb.WriteString(fmt.Sprintf("%s %d\n", p.Category, p.Amount))
			}
			list.SetText(wb.String())
		}

		W.SetContent(container.NewVBox(profile, buttonAdd, dropdown, list))
		return
	}

	label := widget.NewLabel("Вы еще не добавляли доходы")

	W.SetContent(container.NewVBox(profile, label, buttonAdd))
}

func WindowAddIncome() {
	profile := widget.NewLabel("Доходы")
	profile.Alignment = fyne.TextAlignCenter

	entry := widget.NewEntry()
	entry.Keyboard()
	newCategory := widget.NewEntry()
	//newCategory.Hide()
	incomes := db.GetProfit("")
	result := widget.NewLabel("Некорректно введенная информация")
	result.Wrapping = fyne.TextWrapWord
	result.Hide()

	var options = []string{"Добавить категорию"}
	for _, p := range incomes {
		flag := true
		for _, elem := range options {
			if elem == p.Category {
				flag = false
				break
			}
		}
		if flag {
			options = append(options, p.Category)
		}
	}

	dropdown := widget.NewSelect(options, func(elem string) {})
	dropdown.SetSelectedIndex(0)

	buttonAdd := widget.NewButton("Добавить", func() {
		index := options[dropdown.SelectedIndex()]

		amount, err := strconv.Atoi(entry.Text)
		if err != nil {
			result.Show()
			entry.SetText("")
			return
		}

		if newCategory.Hidden {
			status = index
		} else {
			status = newCategory.Text
		}

		db.AddProfit(status, amount)

		result.SetText("Запись добавленна")
		entry.SetText("")
		result.Show()
		WindowIncome()
	})

	dropdown.OnChanged = func(elem string) {
		switch elem {
		default:
			newCategory.Hide()
		case "Добавить категорию":
			fmt.Println("новая категория")
			newCategory.Show()
		}
	}

	content := container.NewVBox(profile, entry, newCategory, dropdown, buttonAdd, result)

	W.SetContent(content)
}
