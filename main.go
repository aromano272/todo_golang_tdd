package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/aromano272/todo_golang_tdd/controllers"
	"gopkg.in/mgo.v2"
)



func main() {
	router := mux.NewRouter()

	tc := controllers.NewTodoController(getSession())

	router.HandleFunc("/get_todos", tc.GetTodos).Methods("GET")
	router.HandleFunc("/get_todo", tc.GetTodo).Methods("GET")
	router.HandleFunc("/create_todo", tc.CreateTodo).Methods("POST")
	router.HandleFunc("/update_todo", tc.UpdateTodo).Methods("PUT")
	router.HandleFunc("/delete_todo", tc.DeleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
