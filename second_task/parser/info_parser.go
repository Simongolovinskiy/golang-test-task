package parser

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

type Influencer struct {
	Rank          int
	Username      string
	Name          string
	Category      string
	Followers     string
	Country       string
	EngAuthentic  string
	EngAverage    string
	FreeReportURL string
}

type HtmlParser struct{}

func (h HtmlParser) GetHtmlInfo(DomHtml *goquery.Document) []Influencer {
	var influencers []Influencer

	tables := DomHtml.Find(".table")

	for i := range tables.Nodes {
		table := tables.Eq(i)

		rows := table.Find(".row")
		rows.Each(func(j int, row *goquery.Selection) {
			influencer := h.ParseRow(row)
			if influencer.Name != "" {
				influencers = append(influencers, influencer)
			}

		})
	}
	return influencers
}

func (h HtmlParser) ParseRow(row *goquery.Selection) Influencer {
	influencer := Influencer{}

	influencer.Rank, _ = strconv.Atoi(row.Find(".rank span").Text())
	influencer.Username = row.Find(".contributor__name-content").Text()
	influencer.Name = row.Find(".contributor__title").Text()
	influencer.Category = row.Find(".category .tag__content").Text()
	influencer.Followers = row.Find(".subscribers").Text()
	influencer.Country = row.Find(".audience").Text()
	influencer.EngAuthentic = row.Find(".authentic").Text()
	influencer.EngAverage = row.Find(".engagement").Text()
	influencer.FreeReportURL, _ = row.Find(".free-report a").Attr("href")

	return influencer
}
