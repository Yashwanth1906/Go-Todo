package main

import (
	"context"
	"github/Yashwanth1906/Go-Todo/pkg/models"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Todo methods that will be called from frontend
func (a *App) AddTask(name, description string) (*models.Task, error) {
	task := &models.Task{
		Name:        name,
		Description: description,
		Status:      models.Pending,
	}
	return models.AddTask(task)
}

func (a *App) GetTasks() []models.Task {
	return models.GetTasks()
}

func (a *App) DeleteTask(id int64) (models.Task, error) {
	return models.DeleteTask(id)
}

func (a *App) GetTaskById(id int64) (models.Task, error) {
	return models.GetTaskById(id)
}

func (a *App) UpdateTaskStatus(id int64) (models.Task, error) {
	return models.UpdateStatusById(id)
}