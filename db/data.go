package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type income struct {
	Category string
	Amount   int
}

type outlay struct {
	Category string
	Amount   int
}

// AddProfit Функция добавления дохода.
// Аргументы:
//
//	category - категория
//	profit - колво денег, что были заработаны
func AddProfit(category string, profit int) string {
	db, err := sql.Open("sqlite3", "db/data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("insert into income (category, amount) values ($1, $2)",
		category, profit)
	if err != nil {
		return "Error"
	}
	result.LastInsertId()
	return "Done"
}

// GetProfit Функция получения информации о деньгах (доход)
// Если category=="", возвращает все позиции в бд
func GetProfit(category string) []income {
	db, err := sql.Open("sqlite3", "db/data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Income")
	if err != nil {
		panic(err)
	}

	if category != "" {
		rows, err = db.Query("SELECT * FROM Income WHERE category = ?", category)
		if err != nil {
			panic(err)
		}
	}

	defer rows.Close()
	incomes := []income{}

	for rows.Next() {
		p := income{}
		err := rows.Scan(&p.Category, &p.Amount)
		if err != nil {
			fmt.Println(err)
			continue
		}
		incomes = append(incomes, p)
	}
	//for _, p := range incomes {
	//	fmt.Println(p.category, p.amount)
	//}
	return incomes
}

func AddOutlay(category string, profit int) string {
	db, err := sql.Open("sqlite3", "db/data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Exec("insert into Outlay (category, amount) values ($1, $2)",
		category, profit)
	if err != nil {
		return "Error"
	}
	result.LastInsertId()
	return "Done"
}

// GetOutlay Функция получения информации о деньгах (доход)
// Если category=="", возвращает все позиции в бд
func GetOutlay(category string) []outlay {
	db, err := sql.Open("sqlite3", "db/data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM Outlay")
	if err != nil {
		panic(err)
	}

	if category != "" {
		rows, err = db.Query("SELECT * FROM Outlay WHERE category = ?", category)
		if err != nil {
			panic(err)
		}
	}

	defer rows.Close()
	outlays := []outlay{}

	for rows.Next() {
		p := outlay{}
		err := rows.Scan(&p.Category, &p.Amount)
		if err != nil {
			fmt.Println(err)
			continue
		}
		outlays = append(outlays, p)
	}
	//for _, p := range outlays {
	//	fmt.Println(p.Category, p.Amount)
	//}
	return outlays
}

func main() {
	//AddProfit("Salary", 999)
	//GetProfit("Salary")
	//AddOutlay("Проезд", 33)
	//AddOutlay("Обучение", 1800)
	GetOutlay("")
}
