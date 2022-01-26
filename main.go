package main

import (
	"log"
	"net/http"

	"github.com/visheshc14/todo/controllers"
	"github.com/visheshc14/todo/models"
	"github.com/gorilla/mux"
)

func InitializeRouters() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/todos", controllers.GetAllTodos).Methods("GET")
	router.HandleFunc("/todo/{taskId}", controllers.GetTodo).Methods("GET")
	router.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todo/{taskId}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{taskId}", controllers.DeleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	models.InitialMigration()
	InitializeRouters()

}
