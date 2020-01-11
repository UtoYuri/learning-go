package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Config struct {
	Dialect string
	Url string
}

var DB *gorm.DB

func Init(config *Config) error {
	db, err := gorm.Open(config.Dialect, config.Url)
	if err != nil {
		return err
	}

	DB = db
	return nil
}
