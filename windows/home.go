package windows

import (
	"awesomeProject2/db"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func Window1() *fyne.Container {
	profile := widget.NewLabel("Главная")
	profile.Alignment = fyne.TextAlignCenter

	incomes := db.GetProfit("")
	outlays := db.GetOutlay("")

	sumIncome := 0
	for _, i := range incomes {
		sumIncome += i.Amount
	}
	sumOutlay := 0
	for _, i := range outlays {
		sumOutlay += i.Amount
	}
	labelSumIncome := widget.NewLabel("Суммарно доходов:\n" + strconv.Itoa(sumIncome))
	labelSumOutlay := widget.NewLabel("Суммарно расходов:\n" + strconv.Itoa(sumOutlay))
	//labelSumIncomeResult := widget.NewLabel(strconv.Itoa(sumIncome))
	//entryResultIncome := widget.NewEntry()
	//entryResultIncome.Disable()
	//entryResultIncome.SetText(strconv.Itoa(sumIncome))

	fmt.Println(sumIncome)

	for _, p := range incomes {
		fmt.Println(p.Category, p.Amount)
	}
	return container.NewVBox(profile, labelSumIncome, labelSumOutlay)
}
