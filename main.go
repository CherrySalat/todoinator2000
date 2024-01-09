package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

const PORT = ":10000"

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var Tasks = []Task{
	Task{
		Id:          int(1),
		Title:       "Buy some bread",
		Description: "1. Go to market\n2.take some bread\n3.Buy it",
	},
	Task{
		Id:          2,
		Title:       "Be cool",
		Description: "1. Put on black glasses 2. Kick some asses 3. Drink beer.",
	},
}

func handleRequests() {
	m := mux.NewRouter().StrictSlash(true)
	m.Handle("/", http.FileServer(http.Dir("./static")))
	m.HandleFunc("/tasks", returnAllTasks)
	m.HandleFunc("/task", createNewTask).Methods("POST")
	m.HandleFunc("/task/{id:[0-9]+}", deleteTask).Methods("DELETE")
	m.HandleFunc("/task/{id:[0-9]+}", returnSingleTask)
	log.Fatal(http.ListenAndServe(":10000", m))
}

func returnSingleTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for _, task := range Tasks {
		if task.Id == id {
			json.NewEncoder(w).Encode(task)
		}
	}
}

func returnAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Tasks)
}

func createNewTask(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var t Task
	json.Unmarshal(reqBody, &t)
	Tasks = append(Tasks, t)
	json.NewEncoder(w).Encode(t)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for index, task := range Tasks {
		if task.Id == id {
			Tasks = append(Tasks[:index], Tasks[index+1:]...)
		}
	}
}

func main() {
	handleRequests()
}
