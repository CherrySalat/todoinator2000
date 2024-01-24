package main

import (
	"encoding/json"
	"github.com/CherrySalat/todoinator2000/model"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const PORT = ":10000"

// TODO Добавить структуру для тела разного рода ошибок и добавить функцию для их вызыва
func main() {
	h := NewTodoHandler(model.LocalData)
	m := mux.NewRouter().StrictSlash(true)
	m.HandleFunc("/todos", h.ListTodo).Methods("GET")
	m.HandleFunc("/todos/{id}", h.GetTodo).Methods("GET")
	m.HandleFunc("/todos", h.CreateTodo).Methods("POST")
	m.HandleFunc("/todos/{id}", h.UpdateTodo).Methods("PUT")
	m.HandleFunc("todos/{id}", h.DeleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(PORT, m))
}

type TodoStore interface {
	Remove(id string) error
	Update(id string, todo model.Todo) error
	List() (map[string]model.Todo, error)
	Get(id string) (model.Todo, error)
	Add(id string, todo model.Todo) error
}

type TodoHandler struct {
	store TodoStore
}

func NewTodoHandler(s TodoStore) *TodoHandler {
	return &TodoHandler{store: s}
}

// TODO Добавить обработку ошибок
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {

	id, _ := GenerateId()
	var todo model.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	h.store.Add(id, todo)
}

func (h *TodoHandler) ListTodo(w http.ResponseWriter, r *http.Request) {
	todos, _ := h.store.List()
	jsonBytes, _ := json.Marshal(todos)
	_ = jsonBytes
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	recipe, _ := h.store.Get(id)
	jsonBytes, _ := json.Marshal(recipe)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		//TODO Добавить обработчик ошибок
		return
	}

	if err := h.store.Update(id, todo); err != nil {
		//TODO Добавить обработчик ошибок
		return
	}

	//TODO Тут тоже обработать
	jsonBytes, _ := json.Marshal(todo)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	//TODO Добавить обработчик ошибок
	h.store.Remove(id)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("its ok"))
}

func GenerateId() (string, error) {
	id, _ := uuid.NewRandom()
	return id.String(), nil
}
