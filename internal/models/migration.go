package models

import (
	"go-module/pkg/database"
)

func AutoMigrate() {
	database.DB.AutoMigrate(&Media{})
}