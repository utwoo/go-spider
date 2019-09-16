package main

import (
	// configuration initialization
	_ "utwoo.com/go-spider/infrastructure/config"

	"utwoo.com/go-spider/infrastructure/db"
)

func main() {
	defer db.DB.Close()
}
