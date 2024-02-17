package utils

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"golang-test-task/second_task/parser"
)

func SaveToExcel(data []parser.Influencer) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Influencers")
	if err != nil {
		fmt.Println("Ошибка при добавлении листа:", err)
		return
	}
	titles := sheet.AddRow()
	titles.AddCell().Value = "Rank"
	titles.AddCell().Value = "Username"
	titles.AddCell().Value = "Name"
	titles.AddCell().Value = "Category"
	titles.AddCell().Value = "Followers"
	titles.AddCell().Value = "Country"
	titles.AddCell().Value = "EngAuthentic"
	titles.AddCell().Value = "EngAverage"
	titles.AddCell().Value = "FreeReportURL"
	for i, influencer := range data {
		row := sheet.AddRow()
		row.AddCell().Value = fmt.Sprintf("%d", i+1)
		row.AddCell().Value = influencer.Username
		row.AddCell().Value = influencer.Name
		row.AddCell().Value = influencer.Category
		row.AddCell().Value = influencer.Followers
		row.AddCell().Value = influencer.Country
		row.AddCell().Value = influencer.EngAuthentic
		row.AddCell().Value = influencer.EngAverage
		row.AddCell().Value = influencer.FreeReportURL
	}
	err = file.Save("output_data.xlsx")
	if err != nil {
		fmt.Println("Ошибка при сохранении файла:", err)
		return
	}

}
