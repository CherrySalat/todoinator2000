package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const PORT = ":10000"

func handleRequests() {
	m := mux.NewRouter().StrictSlash(true)
	m.Handle("/", http.FileServer(http.Dir("./static")))
	m.HandleFunc("/tasks", returnAllTasks)
	m.HandleFunc("/task", createNewTask).Methods("POST")
	m.HandleFunc("/task/{id:[0-9]+}", deleteTask).Methods("DELETE")
	m.HandleFunc("/task/{id:[0-9]+}", returnSingleTask)
	log.Fatal(http.ListenAndServe(PORT, m))
}
