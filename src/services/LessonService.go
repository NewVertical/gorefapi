package services

import (
	"apiref/core"
	"apiref/src/models"
	"database/sql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CommonService interface {
}

type LessonService struct {
	db *gorm.DB
}

func (l LessonService) New() LessonService {
	db, _ := gorm.Open(postgres.Open(core.ConnectInfo{}.New().ConnectionString()), &gorm.Config{})
	l.db = db
	return l
}
func (l LessonService) Create(lesson *models.Lesson) LessonService {
	l.db.Create(&lesson)
	return l
}
func (l LessonService) Update(lesson *models.Lesson) LessonService {
	l.db.Save(&lesson)
	return l
}
func (l LessonService) Delete(id int) LessonService {
	l.db.Delete(&models.Lesson{}, id)
	return l
}
func (l LessonService) GetList() ([]models.Lesson, error) {
	var lessons []models.Lesson
	l.db.Find(&lessons)
	return lessons, nil
}
func (l LessonService) RawGetList() ([]any, error) {
	return core.ConnectInfo{User: "postgres", Pass: "password", Database: "refdb"}.ExecuteSelect("SELECT * FROM lessons LIMIT $1 OFFSET $2",
		func(rows *sql.Rows) ([]any, error) {
			var result []any
			for rows.Next() {
				var p models.Lesson
				if err := rows.Scan(&p.ID, &p.Author, &p.Title); err != nil {
					return nil, err
				}
				result = append(result, p)
			}
			return result, nil
		})
}
