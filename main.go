package main

import (
	"github.com/aromano272/todo_golang_tdd/controllers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"github.com/aromano272/todo_golang_tdd/data"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	todoDAO := data.NewTodoDAO(getSession())
	tc := controllers.NewTodoController(todoDAO)

	router.HandleFunc("/todos", tc.GetTodos).Methods("GET")
	router.HandleFunc("/todos/{id}", tc.GetTodo).Methods("GET")
	router.HandleFunc("/todos", tc.CreateTodo).Methods("POST")
	router.HandleFunc("/todos", tc.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos", tc.DeleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
