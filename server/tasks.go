// tasks.go

package main

import (
	"fmt"
	"sync"
)

// Task representa una tarea en la lista de tareas.
type Task struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

// TaskManager maneja la gestión de tareas.
type TaskManager struct {
	tasks      []*Task
	mutex      sync.Mutex
	lastTaskID int
}

// NewTaskManager crea una nueva instancia de TaskManager.
func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

// AddTask agrega una nueva tarea a la lista de tareas.
func (tm *TaskManager) AddTask(name string) *Task {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	tm.lastTaskID++
	task := &Task{
		ID:     fmt.Sprintf("%d", tm.lastTaskID),
		Name:   name,
		Status: "pending",
	}
	tm.tasks = append(tm.tasks, task)
	return task
}

// GetTasks devuelve la lista de todas las tareas.
func (tm *TaskManager) GetTasks() []*Task {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	return tm.tasks
}

// UpdateTaskStatus actualiza el estado de una tarea específica.
func (tm *TaskManager) UpdateTaskStatus(taskID string, status string) error {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	for _, task := range tm.tasks {
		if task.ID == taskID {
			task.Status = status
			return nil
		}
	}
	return fmt.Errorf("task with ID %s not found", taskID)
}

// DeleteTask elimina una tarea de la lista de tareas.
func (tm *TaskManager) DeleteTask(taskID string) error {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	for i, task := range tm.tasks {
		if task.ID == taskID {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %s not found", taskID)
}
