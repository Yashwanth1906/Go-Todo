package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Status int

const (
	Pending Status = iota
	Completed
)

type Task struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

func addTask(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// fmt.Println(params["name"]) // This type of usage is used for getting data from the URL for get Method Not Post Method.
	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTask.ID = int64(rand.Intn(100000))
	newTask.Status = Pending
	tasks[newTask.ID] = newTask
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getTaskById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	fmt.Println("Params : ", idStr)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	task, exists := tasks[id]
	if !exists {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func deleteTaskById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task Id", http.StatusBadRequest)
		return
	}
	_, exists := tasks[id]
	delete(tasks, id)
	if !exists {
		http.Error(w, "No Task is found with the id", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	fmt.Println(task)
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	already, exists := tasks[id]
	if !exists {
		http.Error(w, "Task not found with the Id", http.StatusNotFound)
		return
	}
	if task.Name != "" {
		already.Name = task.Name
	}
	if task.Description != "" {
		already.Description = task.Description
	}
	tasks[id] = already
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Task Updated successfully"})
}

func updateTaskStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	already, exists := tasks[id]
	if !exists {
		http.Error(w, "Task not found with the Id", http.StatusNotFound)
		return
	}
	if already.Status == Pending {
		already.Status = Completed
	} else {
		already.Status = Pending
	}
	tasks[id] = already
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Task Marked successfully"})
}

var tasks map[int64]Task = make(map[int64]Task)

func main() {
	r := mux.NewRouter()
	fmt.Println(tasks)
	tasks[1] = Task{
		Name:        "Do Homework",
		Description: "O",
		Status:      Pending,
	}
	r.HandleFunc("/taskadd", addTask).Methods("POST")
	r.HandleFunc("/task/{id}", getTaskById).Methods("GET")
	r.HandleFunc("/task/{id}", deleteTaskById).Methods("DELETE")
	r.HandleFunc("/task/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/tasks", getAllTasks).Methods("GET")
	r.HandleFunc("/task/updatestatus/{id}", updateTaskStatus).Methods("PUT")

	fmt.Printf("Starting Server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
