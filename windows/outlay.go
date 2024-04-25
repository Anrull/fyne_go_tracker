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

var status = "default"

func WindowOutlay() {
	profile := widget.NewLabel("Расходы")
	profile.Alignment = fyne.TextAlignCenter

	outlays := db.GetOutlay("")

	var options = []string{"Все позиции"}
	for _, p := range outlays {
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

	buttonAdd := widget.NewButton("Добавить", func() { WindowAddOutlay() })

	//toolbar := widget.NewToolbar(
	//	//widget.NewToolbarSpacer(),
	//	widget.NewToolbarAction(theme.ContentAddIcon(), func() { fmt.Println("1") }),
	//	widget.NewToolbarAction(theme.HelpIcon(), func() {
	//		log.Println("Display help")
	//	}),
	//)

	//
	//content := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("Content"))

	if len(outlays) > 0 {
		dropdown := widget.NewSelect(options, func(elem string) { status = elem })
		dropdown.SetSelectedIndex(0)

		list := widget.NewLabel("")

		wb := &strings.Builder{}
		for _, p := range outlays {
			wb.WriteString(fmt.Sprintf("%s %d\n", p.Category, p.Amount))
		}
		list.SetText(wb.String())

		dropdown.OnChanged = func(elem string) {
			if elem == "Все позиции" {
				elem = ""
			}
			outlays := db.GetOutlay(elem)
			wb := &strings.Builder{}
			for _, p := range outlays {
				wb.WriteString(fmt.Sprintf("%s %d\n", p.Category, p.Amount))
			}
			list.SetText(wb.String())
		}

		W.SetContent(container.NewVBox(profile, buttonAdd, dropdown, list))
		return
	}

	label := widget.NewLabel("Вы еще не добавляли расходы")

	W.SetContent(container.NewVBox(profile, label, buttonAdd))
}

func WindowAddOutlay() {
	profile := widget.NewLabel("Расходы")
	profile.Alignment = fyne.TextAlignCenter

	entry := widget.NewEntry()
	entry.Keyboard()
	newCategory := widget.NewEntry()
	//newCategory.Hide()
	outlays := db.GetOutlay("")
	result := widget.NewLabel("Некорректно введенная информация")
	result.Wrapping = fyne.TextWrapWord
	result.Hide()

	var options = []string{"Добавить категорию"}
	for _, p := range outlays {
		options = append(options, p.Category)
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

		db.AddOutlay(status, amount)

		result.SetText("Запись добавленна")
		entry.SetText("")
		result.Show()
		WindowOutlay()
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
	//if number, err := strconv.Atoi(entry.Text); err != nil {
	//
	//}
}
