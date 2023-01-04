package routes

import (
	"restAPI/CRUD/pkg/controllers"

	"github.com/gorilla/mux"
)

func addMainHandlers(r *mux.Router) {
	r.HandleFunc("/movies", controllers.GetAllMoviesHandler).Methods("GET")
	r.HandleFunc("/movie", controllers.CreateMovieHandler).Methods("POST")
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	addMainHandlers(r)
	return r
}
