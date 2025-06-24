package models

import (
	"fmt"

	"github.com/Yashwanth1906/Go-Todo/backend/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Status int

const (
	Pending Status = iota
	Completed
)

type Task struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	if err := db.AutoMigrate(&Task{}); err != nil {
		panic("Failed to auto-migrate Task: " + err.Error())
	}
}

func AddTask(task *Task) (*Task, error) {
	result := db.Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func GetTasks() []Task {
	var tasks []Task
	db.Find(&tasks)
	return tasks
}

func DeleteTask(id int64) (Task, error) {
	var task Task

	result := db.First(&task, id)
	if result.Error != nil {
		fmt.Println("Task not found:", result.Error)
		return Task{}, result.Error
	}

	if err := db.Delete(&task).Error; err != nil {
		fmt.Println("Delete failed:", err)
		return Task{}, err
	}

	return task, nil
}

func GetTaskById(id int64) (Task, error) {
	var task Task
	result := db.First(&task, id)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func UpdateStatusById(id int64) (Task, error) {
	var task Task
	result := db.First(&task, id)
	if result.Error != nil {
		return Task{}, result.Error
	}
	if task.Status == Completed {
		task.Status = Pending
	} else {
		task.Status = Completed
	}
	if err := db.Save(&task).Error; err != nil {
		fmt.Println("Delete failed:", err)
		return Task{}, err
	}
	return task, nil
}
