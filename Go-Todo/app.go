package main

import (
	"context"
	"github/Yashwanth1906/Go-Todo/pkg/config"
	"github/Yashwanth1906/Go-Todo/pkg/models"
	"log"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	log.Println("Creating new App instance")
	return &App{}
}

// This is not even working I don't know why..
func (a *App) startup(ctx context.Context) {
	log.Println("App startup called")
	a.ctx = ctx

	log.Println("Initializing database connection...")
	config.Connect()
	log.Println("Initializing models...")

	log.Println("App startup completed successfully")
}

// Todo methods that will be called from frontend
func (a *App) AddTask(name, description string) (*models.Task, error) {
	log.Printf("AddTask called with name: %s, description: %s", name, description)

	task := &models.Task{
		Name:        name,
		Description: description,
		Status:      models.Pending,
	}

	result, err := models.AddTask(task)
	if err != nil {
		log.Printf("AddTask error: %v", err)
		return nil, err
	}

	log.Printf("AddTask success: %+v", result)
	return result, nil
}

func (a *App) GetTasks() []models.Task {
	log.Println("GetTasks called")

	tasks := models.GetTasks()
	log.Printf("GetTasks returned %d tasks", len(tasks))

	for i, task := range tasks {
		log.Printf("Task %d: %+v", i, task)
	}

	return tasks
}

func (a *App) DeleteTask(id int64) (models.Task, error) {
	log.Printf("DeleteTask called with id: %d", id)

	result, err := models.DeleteTask(id)
	if err != nil {
		log.Printf("DeleteTask error: %v", err)
		return models.Task{}, err
	}

	log.Printf("DeleteTask success: %+v", result)
	return result, nil
}

func (a *App) GetTaskById(id int64) (models.Task, error) {
	log.Printf("GetTaskById called with id: %d", id)

	result, err := models.GetTaskById(id)
	if err != nil {
		log.Printf("GetTaskById error: %v", err)
		return models.Task{}, err
	}

	log.Printf("GetTaskById success: %+v", result)
	return result, nil
}

func (a *App) UpdateTaskStatus(id int64) (models.Task, error) {
	log.Printf("UpdateTaskStatus called with id: %d", id)

	result, err := models.UpdateStatusById(id)
	if err != nil {
		log.Printf("UpdateTaskStatus error: %v", err)
		return models.Task{}, err
	}

	log.Printf("UpdateTaskStatus success: %+v", result)
	return result, nil
}
