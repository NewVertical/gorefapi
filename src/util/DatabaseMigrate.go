package util

import (
	"apiref/core"
	"apiref/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseMigrate() {
	db, err := gorm.Open(postgres.Open(core.ConnectInfo{}.New().ConnectionString()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	model_err := db.AutoMigrate(&models.Lesson{})
	if model_err != nil {
		return
	}

}
