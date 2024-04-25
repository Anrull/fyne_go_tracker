package main

import (
	"awesomeProject2/lexicon"
	"awesomeProject2/windows"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var myApp = windows.MyApp
var w = windows.W

func showHome() {
	fmt.Println("Home")
	w.SetContent(windows.Window1())
}

func showIncome() {
	fmt.Println("Income")
	windows.WindowIncome()
}

func showOutlay() {
	fmt.Println("Outlay")
	windows.WindowOutlay()
}

func showInfoAboutDeveloper() {
	profile := widget.NewLabel("О разработчиках")
	profile.Alignment = fyne.TextAlignCenter

	text := widget.NewLabel(lexicon.InfoAboutDeveloper)
	text.Wrapping = fyne.TextWrapWord
	content := container.NewVBox(profile, text)
	w.SetContent(content)
}

func showInfoAboutTechnology() {
	profile := widget.NewLabel("О технологиях")
	profile.Alignment = fyne.TextAlignCenter

	text := widget.NewLabel(lexicon.InfoAboutTechnology)
	text.Wrapping = fyne.TextWrapWord
	content := container.NewVBox(profile, text)
	w.SetContent(content)
}

func showInfoAboutProject() {
	profile := widget.NewLabel("О проекте")
	profile.Alignment = fyne.TextAlignCenter

	text := widget.NewLabel(lexicon.InfoAboutProject)
	text.Wrapping = fyne.TextWrapWord
	content := container.NewVBox(profile, text)
	w.SetContent(content)
}

func main() {
	w.Resize(fyne.Size{Width: 275, Height: 375})

	menu := fyne.NewMainMenu(
		fyne.NewMenu("Основное",
			fyne.NewMenuItem("Главная", func() { showHome() }),
			fyne.NewMenuItem("Расходы", func() { showOutlay() }),
			fyne.NewMenuItem("Доходы", func() { showIncome() }),
		),
		fyne.NewMenu("Информация",
			fyne.NewMenuItem("О разработчиках", func() { showInfoAboutDeveloper() }),
			fyne.NewMenuItem("О технологиях", func() { showInfoAboutTechnology() }),
			fyne.NewMenuItem("О проекте", func() { showInfoAboutProject() }),
		),
	)

	w.SetMainMenu(menu)
	w.SetContent(windows.Window1())
	w.ShowAndRun()
}

//toolbar := widget.NewToolbar(
//	widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
//		log.Println("New document")
//	}),
//	widget.NewToolbarSeparator(),
//	widget.NewToolbarAction(theme.ContentAddIcon(), func() { fmt.Println("1") }),
//	widget.NewToolbarAction(theme.ContentCopyIcon(), func() { fmt.Println(2) }),
//	widget.NewToolbarAction(theme.ContentPasteIcon(), func() { fmt.Println(3) }),
//	widget.NewToolbarSpacer(),
//	widget.NewToolbarAction(theme.HelpIcon(), func() {
//		log.Println("Display help")
//	}),
//)
//
//content := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("Content"))

//var homeButton = widget.NewButton("Главная", showHome)
//var incomeButton = widget.NewButton("Доходы", func() { fmt.Println("income") })
//var outlayButton = widget.NewButton("Расходы", func() { fmt.Println("Outlay") })
//
//var menuButtons = container.NewHBox(homeButton, incomeButton, outlayButton)
