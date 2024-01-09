package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

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
