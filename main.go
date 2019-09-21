package main

import (
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

	//InitialTagProcessInfo()
}

// InitialTagProcessInfo initialize the spider process information by tag
func InitialTagProcessInfo() {
	urlTags := config.BookTagURL
	docTags, _ := download.Downloader(urlTags)
	domain.TagProcessPhase(docTags)
}
