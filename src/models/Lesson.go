package models

import "gorm.io/gorm"

type Lesson struct {
	gorm.Model
	ID        int
	Author    string
	Title     string
	Publisher string
}
