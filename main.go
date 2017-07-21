package main

import (
	"github.com/aromano272/todo_golang_tdd/controllers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
	"github.com/aromano272/todo_golang_tdd/data"
	"github.com/aromano272/todo_golang_tdd/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	th := handlers.NewTodoHandler(
		controllers.NewTodoController(
			data.NewTodoDAO(getSession()),
		),
	)

	router.Handle("/todos", Wrap(th.ReadAllTodos)).Methods("GET")
	router.Handle("/todos/{id}", Wrap(th.ReadTodo)).Methods("GET")
	router.Handle("/todos", Wrap(th.CreateTodo)).Methods("POST")
	router.Handle("/todos", Wrap(th.UpdateTodo)).Methods("PUT")
	router.Handle("/todos/{id}", Wrap(th.DeleteTodo)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func Wrap(h http.HandlerFunc) http.Handler {
	return &wrapper{h}
}

type wrapper struct {
	handler http.Handler
}

func (wrpr *wrapper) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	// ↑↑↑ before each request ↑↑↑
	wrpr.handler.ServeHTTP(res, req)
	// ↓↓↓ after each request  ↓↓↓

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
