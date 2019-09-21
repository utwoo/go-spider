package postgres

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	// postgresql provider for gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// InitializeDatabase initialize database and return it
func InitializeDatabase() *gorm.DB {
	var (
		dbName   string
		port     string
		user     string
		sslMode  string
		password string
		host     string
	)

	dbName = viper.GetString("db.dbname")
	port = viper.GetString("db.port")
	user = viper.GetString("db.user")
	sslMode = viper.GetString("db.sslmode")
	password = viper.GetString("db.password")
	host = viper.GetString("db.host")

	connectString := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s password=%s", host, user, dbName, port, sslMode, password)
	fmt.Println(connectString)

	connect, err := gorm.Open(
		"postgres",
		connectString,
	)
	if err != nil {
		fmt.Println(err)
		panic("connect postgres failed")
	}

	connect.LogMode(true)
	fmt.Println("Login postgres database success!")
	return connect
}
