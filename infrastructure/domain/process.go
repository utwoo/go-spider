package domain

import (
	"github.com/PuerkitoBio/goquery"
	"utwoo.com/go-spider/infrastructure/adapter"
	"utwoo.com/go-spider/infrastructure/db"
	"utwoo.com/go-spider/infrastructure/model"
)

func TagProcessPhase(doc *goquery.Document) {
	doc.Find(".tagCol td").Each(func(i int, selection *goquery.Selection) {
		tagName := adapter.ToString(selection.Find("a").Text())
		tagURL, _ := selection.Find("a").Attr("href")

		processes := model.TagProcess{
			TagName:    tagName,
			TagURL:     tagURL,
			StartIndex: 0,
			Finished:   "N",
		}

		db.DB.Save(&processes)
	})
}
