package interfaces

import "github.com/PuerkitoBio/goquery"
import "golang-test-task/second_task/parser"

type ApiParserInt interface {
	GetData() *goquery.Document
}

type HtmlParserInt interface {
	GetHtmlInfo(DomHtml *goquery.Document) []parser.Influencer
	ParseRow(row *goquery.Selection) parser.Influencer
}
