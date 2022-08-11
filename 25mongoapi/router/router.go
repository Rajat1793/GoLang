package router

import (
	"mongoapi/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	// routing
	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	r.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT") //with the route we have to pass the {id} which is same as params id what we have passed in other function (both should be same)
	r.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	r.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovie).Methods("DELETE")

	return r
}
