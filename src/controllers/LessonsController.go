package controllers

import (
	"apiref/src/controllers/utils"
	"apiref/src/models"
	"apiref/src/services"
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

type LessonsController struct {
	Router *mux.Router
	DB     *sql.DB
}

func (l LessonsController) GetRouter() *mux.Router {
	return l.Router
}

func (l LessonsController) New() LessonsController {
	l.Router.HandleFunc("", l.list).Methods("GET")
	l.Router.HandleFunc("/", l.list).Methods("GET")
	l.Router.HandleFunc("/{id}", l.get).Methods("GET")
	l.Router.HandleFunc("", l.create).Methods("POST")
	l.Router.HandleFunc("/", l.create).Methods("POST")
	l.Router.HandleFunc("/{id}", l.update).Methods("PUT")
	l.Router.HandleFunc("/{id}", l.delete).Methods("DELETE")
	return l
}

// ToDo: Still need to work out the paging from the server.
func (l LessonsController) list(w http.ResponseWriter, r *http.Request) {
	result, err := services.LessonService{}.New().GetList()
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error Getting Rows")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, result)
}
func (l LessonsController) get(w http.ResponseWriter, r *http.Request) {
	utils.GetOneRequest(w, r, func(id int) {
		var lesson models.Lesson
		lesson.ID = id
		utils.RespondWithJSON(w, http.StatusOK, lesson)
	})
}

func (l LessonsController) create(w http.ResponseWriter, r *http.Request) {
	utils.CreateRequest[models.Lesson](w, r, func(lesson models.Lesson) {
		services.LessonService{}.New().Create(&lesson)
		utils.RespondWithJSON(w, http.StatusCreated, lesson)
	})
}

func (l LessonsController) update(w http.ResponseWriter, r *http.Request) {
	utils.UpdateRequest[models.Lesson](w, r, func(id int, lesson models.Lesson) {
		services.LessonService{}.New().Update(&lesson)
		utils.RespondWithJSON(w, http.StatusOK, lesson)
	})
}

func (l LessonsController) delete(w http.ResponseWriter, r *http.Request) {
	utils.DeleteRequest(w, r, func(id int) {
		services.LessonService{}.New().Delete(id)
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	})
}
