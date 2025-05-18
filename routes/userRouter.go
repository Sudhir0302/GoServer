package routes

import (
	"github.com/Sudhir0302/GoServer/controller"
	"github.com/gorilla/mux"
)

func UserRoute() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/test",controller.Greet)

	return router
}