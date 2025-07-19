package routes

import (
	"encoding/json"
	"mongodbapi/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {   //it's like anonymous function
		json.NewEncoder(w).Encode("hello from server")
	})

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/delete/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deletemovies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
