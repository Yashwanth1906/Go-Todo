package routes

import (
	"github/Yashwanth1906/Go-Todo/pkg/controllers"

	"github.com/gorilla/mux"
)

var Routes = func(router *mux.Router) {
	router.HandleFunc("/taskadd", controllers.AddTask).Methods("POST")
	router.HandleFunc("/task/{id}", controllers.GetTaskById).Methods("GET")
	router.HandleFunc("/task/{id}", controllers.DeleteTaskById).Methods("DELETE")
	router.HandleFunc("/task/{id}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks", controllers.GetAllTasks).Methods("GET")
	router.HandleFunc("/task/updatestatus/{id}", controllers.UpdateTaskStatus).Methods("PUT")
}
