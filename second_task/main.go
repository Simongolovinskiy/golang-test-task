package main

import (
	"golang-test-task/second_task/api"
	"golang-test-task/second_task/interfaces"
	"golang-test-task/second_task/parser"
	"golang-test-task/second_task/utils"
)

func run() {
	var Apiparser interfaces.ApiParserInt = api.ApiParser{}
	var myParser interfaces.HtmlParserInt = parser.HtmlParser{}
	parsedData := myParser.GetHtmlInfo(Apiparser.GetData())
	utils.SaveToExcel(parsedData)
}

func main() {
	run()
}
