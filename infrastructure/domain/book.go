package domain

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
	"strings"
	"time"
	"utwoo.com/go-spider/infrastructure/adapter"
	"utwoo.com/go-spider/infrastructure/db"
	"utwoo.com/go-spider/infrastructure/download"
	"utwoo.com/go-spider/infrastructure/model"
)

func SaveBooksPerTagTypePage(doc *goquery.Document) int {
	result := doc.Find("#subject_list .subject-list .subject-item")

	if result.Length() > 0 {
		result.Each(func(i int, selection *goquery.Selection) {
			bookURL := selection.Find(".info a").AttrOr("href", "")
			// Save book information in book page
			document, _ := download.Downloader(bookURL)
			sid := regexp.MustCompile(`[\d]+`).FindString(bookURL)
			SaveBookInfo(sid, document)
		})
	}

	return result.Length()
}

func SaveBookInfo(sid string, doc *goquery.Document) {
	bookInfo := model.Book{SID: sid}
	db.DB.Where(&model.Book{SID: sid}).Find(&bookInfo)

	// Book Title
	bookInfo.Title = adapter.ToString(doc.Find("#wrapper span[property]").First().Text())
	// Book Image URL
	bookInfo.ImageURL = adapter.ToString(doc.Find(".nbg").AttrOr("href", ""))
	// Book Rate
	bookInfo.Rate, _ = strconv.ParseFloat(adapter.ToString(doc.Find(".rating_num").Text()), 64)
	// Book Intro
	bookInfo.Intro, _ = doc.Find(".intro").Html()
	// Book Information
	doc.Find("div#info span.pl").Each(func(i int, selection *goquery.Selection) {
		property := strings.TrimRight(adapter.ToString(selection.Text()), ":")
		switch property {
		case "作者":
			bookInfo.Author = adapter.ToString(selection.NextFiltered("a").Text())
			bookInfo.Author = strings.ReplaceAll(bookInfo.Author, " ", "")
		case "出版社":
			bookInfo.Publisher = adapter.ToString(selection.Nodes[0].NextSibling.Data)
		case "出品方":
			bookInfo.Producer = adapter.ToString(selection.NextFiltered("a").Text())
		case "副标题":
			bookInfo.SubTitle = adapter.ToString(selection.Nodes[0].NextSibling.Data)
		case "原作名":
			bookInfo.OriginTitle = adapter.ToString(selection.Nodes[0].NextSibling.Data)
		case "译者":
			bookInfo.Translator = adapter.ToString(selection.NextFiltered("a").Text())
		case "出版年":
			bookPublished := adapter.ToString(selection.Nodes[0].NextSibling.Data)
			bookPublished = strings.ReplaceAll(bookPublished, "年", "-")
			bookPublished = strings.ReplaceAll(bookPublished, "月", "-")
			bookPublished = strings.ReplaceAll(bookPublished, "日", "")
			bookPublished = strings.Trim(bookPublished, "-")
			if strings.Count(bookPublished, "-") < 2 {
				bookPublished += "-1"
			}
			bookInfo.Published, _ = time.Parse("2006-1-1", bookPublished)
		case "页数":
			bookInfo.PageNumber, _ = strconv.Atoi(adapter.ToString(selection.Nodes[0].NextSibling.Data))
		case "定价":
			price := regexp.MustCompile(`[0-9.]+`).FindString(adapter.ToString(selection.Nodes[0].NextSibling.Data))
			bookInfo.Price, _ = strconv.ParseFloat(price, 64)
		case "装帧":
			bookInfo.CoverType = adapter.ToString(selection.Nodes[0].NextSibling.Data)
		case "丛书":
			bookInfo.Series = adapter.ToString(selection.NextFiltered("a").Text())
		case "ISBN":
			bookInfo.ISBN = adapter.ToString(selection.Nodes[0].NextSibling.Data)
		}
	})

	db.DB.Save(&bookInfo)
}
