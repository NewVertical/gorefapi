package utils

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		return
	}
}

func GetOneRequest(w http.ResponseWriter, r *http.Request, handler func(id int)) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)
	handler(id)
}
func UpdateRequest[K any](w http.ResponseWriter, r *http.Request, handler func(id int, target K)) {
	var source K
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&source); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)
	handler(id, source)
}
func CreateRequest[K any](w http.ResponseWriter, r *http.Request, handler func(target K)) {
	var source K
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&source); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer closeBody(r.Body)
	handler(source)
}
func closeBody(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {

	}
}
func DeleteRequest(w http.ResponseWriter, r *http.Request, handler func(id int)) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(r.Body)
	handler(id)
}
