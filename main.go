package main

import (
	// configuration initialization
	_ "utwoo.com/go-spider/infrastructure/config"
	"utwoo.com/go-spider/infrastructure/db"
	"utwoo.com/go-spider/infrastructure/model"
)

func main() {
	db.DB.AutoMigrate(&model.Book{})
	defer db.DB.Close()
}
