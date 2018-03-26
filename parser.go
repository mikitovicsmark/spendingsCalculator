package main

import (
	"log"

	"github.com/tealeg/xlsx"
)

type Spending struct {
	Location     string `json:"location"`
	SpendingType string `json:"type"`
	Value        int    `json:"value"`
}

type Day struct {
	Weekday   string     `json:"weekday"`
	Spendings []Spending `json:"spending"`
}

func parse(excelFileName string) (days []Day) {
	var count = 0
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatal(err)
	}
	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			for columnIndex, cell := range row.Cells {
				if cell.String() != "" && rowIndex != 0 {
					switch columnIndex {
					case 0:
						count++
						days = append(days, Day{Weekday: cell.String()})
					case 1:
						var cellValue, _ = cell.Int()
						days[count-1].Spendings = append(days[count-1].Spendings, Spending{Value: cellValue})
					case 2:
						var spendingsCount = len(days[count-1].Spendings) - 1
						days[count-1].Spendings[spendingsCount].Location = cell.String()
					case 3:
						var spendingsCount = len(days[count-1].Spendings) - 1
						days[count-1].Spendings[spendingsCount].SpendingType = cell.String()
					}
				}
			}
		}
	}
	return
}
