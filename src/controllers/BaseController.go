package controllers

import "github.com/gorilla/mux"

type IBaseController interface {
	New() LessonsController
	GetRouter() *mux.Router
}
