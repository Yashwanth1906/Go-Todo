package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Yashwanth1906/Go-Todo/pkg/models"
	"github.com/gorilla/mux"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// fmt.Println(params["name"]) // This type of usage is used for getting data from the URL for get Method Not Post Method.
	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTask.Status = models.Pending
	task := models.AddTask(&newTask)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// idStr := params["id"]
	// fmt.Println("Params : ", idStr)
	// id, err := strconv.ParseInt(idStr, 10, 64)
	// if err != nil {
	// 	http.Error(w, "Invalid task ID", http.StatusBadRequest)
	// 	return
	// }
	// task, exists := tasks[id]
	// if !exists {
	// 	http.Error(w, "Task not found", http.StatusNotFound)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(task)
}

func DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	params := mux.Vars(r)
	idStr := params["id"]
	fmt.Println("Came inside Deletetask")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid task Id", http.StatusBadRequest)
		return
	}
	task := models.DeleteTask(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks := models.GetTask()
	json.NewEncoder(w).Encode(tasks)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// var task models.Task
	// json.NewDecoder(r.Body).Decode(&task)
	// fmt.Println(task)
	// params := mux.Vars(r)
	// idStr := params["id"]
	// id, err := strconv.ParseInt(idStr, 10, 64)
	// if err != nil {
	// 	http.Error(w, "Invalid ID", http.StatusBadRequest)
	// 	return
	// }
	// already, exists := tasks[id]
	// if !exists {
	// 	http.Error(w, "Task not found with the Id", http.StatusNotFound)
	// 	return
	// }
	// if task.Name != "" {
	// 	already.Name = task.Name
	// }
	// if task.Description != "" {
	// 	already.Description = task.Description
	// }
	// tasks[id] = already
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(map[string]string{"message": "Task Updated successfully"})
}

func UpdateTaskStatus(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// idStr := params["id"]
	// id, err := strconv.ParseInt(idStr, 10, 64)
	// if err != nil {
	// 	http.Error(w, "Invalid ID", http.StatusBadRequest)
	// 	return
	// }
	// already, exists := tasks[id]
	// if !exists {
	// 	http.Error(w, "Task not found with the Id", http.StatusNotFound)
	// 	return
	// }
	// if already.Status == models.Pending {
	// 	already.Status = models.Completed
	// } else {
	// 	already.Status = models.Pending
	// }
	// tasks[id] = already
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(map[string]string{"message": "Task Marked successfully"})
}
