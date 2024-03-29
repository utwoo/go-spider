package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	// BookRootURL presents the root url of DouBan books
	BookRootURL = "https://book.douban.com"
	// BookTagByCloudURL presents the tags of books
	BookTagByCloudURL = fmt.Sprintf("%s/tag/?view=cloud", BookRootURL)
	// BookTagByType presents the tags of books
	BookTagByTypeURL = fmt.Sprintf("%s/tag/?view=type", BookRootURL)
)

func init() {
	viper.AddConfigPath("./infrastructure/config")
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}
	fmt.Println("read config file successfully")
}
