package db

import (
	"github.com/jinzhu/gorm"
	"utwoo.com/go-spider/infrastructure/db/postgres"
)

// DB presents the database instance
var DB *gorm.DB

func init() {
	DB = postgres.InitializeDatabase()
}
