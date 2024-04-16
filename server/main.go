// main.go

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// Crear una nueva instancia de TaskManager
	taskManager := NewTaskManager()

	// Definir los manejadores para las rutas
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handleAddTask(w, r, taskManager)
		case http.MethodGet:
			handleGetTasks(w, taskManager)
		default:
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			handleUpdateTaskStatus(w, r, taskManager)
		} else if r.Method == http.MethodDelete {
			handleDeleteTask(w, r, taskManager)
		} else {
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})

	// Iniciar el servidor en el puerto 8080
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleAddTask(w http.ResponseWriter, r *http.Request, taskManager *TaskManager) {
	var task struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if task.Name == "" {
		respondWithError(w, http.StatusBadRequest, "Task name is required")
		return
	}

	newTask := taskManager.AddTask(task.Name)
	respondWithJSON(w, http.StatusCreated, newTask)
}

func handleGetTasks(w http.ResponseWriter, taskManager *TaskManager) {
	tasks := taskManager.GetTasks()
	respondWithJSON(w, http.StatusOK, tasks)
}

func handleUpdateTaskStatus(w http.ResponseWriter, r *http.Request, taskManager *TaskManager) {
	taskID := r.URL.Path[len("/tasks/"):]
	var status struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&status); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := taskManager.UpdateTaskStatus(taskID, status.Status); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Task status updated successfully"})
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request, taskManager *TaskManager) {
	taskID := r.URL.Path[len("/tasks/"):]
	if err := taskManager.DeleteTask(taskID); err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
