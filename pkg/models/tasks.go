package models

import (
	"fmt"

	"github.com/Yashwanth1906/Go-Todo/pkg/config"
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

func AddTask(task *Task) *Task {
	db.Create(task)
	return task
}

func GetTask() []Task {
	var Tasks []Task
	db.Find(&Tasks)
	return Tasks
}

func DeleteTask(Id int64) Task {
	fmt.Println("Id : ", Id)
	var task Task
	db.Where("ID=?", Id).Find(&task)
	return task
}
