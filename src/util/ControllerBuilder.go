package util

import (
	"apiref/src/controllers"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"net/http"
)

type ControllerBuilder struct {
}

func (c ControllerBuilder) Build() *mux.Router {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	ApplyTokenAuth(controllers.LessonsController{Router: router.PathPrefix("/v1/lessons").Subrouter()}.New())
	return router
}
func ApplyTokenAuth(controller controllers.IBaseController) controllers.IBaseController {
	controller.GetRouter().Use(SetCORSHeader)
	controller.GetRouter().Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	//controller.GetRouter().Use(RequireTokenAuthentication)
	return controller
}
func SetCORSHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
func RequireTokenAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("secret"), nil
		})
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		context.WithValue(r.Context(), "decoded", token.Claims)

		next.ServeHTTP(w, r)
	})
}
