package main

import (
	"fmt"
	// configuration initialization
	"utwoo.com/go-spider/infrastructure/config"
	"utwoo.com/go-spider/infrastructure/db"
	"utwoo.com/go-spider/infrastructure/domain"
	"utwoo.com/go-spider/infrastructure/download"
	"utwoo.com/go-spider/infrastructure/model"
)

func main() {
	// Migrate Database
	db.DB.AutoMigrate(
		&model.Book{},
		&model.TagProcess{},
	)
	defer db.DB.Close()

	//InitializeTagProcessInfo()

	PushBooks()
}

// InitializeTagProcessInfo initialize the spider process information by tag
func InitializeTagProcessInfo() {
	urlTags := config.BookTagByTypeURL
	document, _ := download.Downloader(urlTags)
	domain.TagProcessPhase(document)
}

func PushBooks() {
	tagProcess := model.TagProcess{}

	for {
		// Get incomplete tag
		db.DB.Where("finished = 'N'").First(&tagProcess)
		// Get books information in tag page
		tagURL := fmt.Sprintf("%s%s?start=%d&type=T", config.BookRootURL, tagProcess.TagURL, tagProcess.StartIndex)
		document, _ := download.Downloader(tagURL)
		result := domain.SaveBooksPerTagTypePage(document)

		if result == 0 {
			tagProcess.Finished = "Y"
		}

		tagProcess.StartIndex += result
		db.DB.Save(&tagProcess)
	}
}
