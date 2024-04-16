// main.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func main() {
	// Realizar la solicitud HTTP para obtener todas las tareas
	response, err := http.Get("http://localhost:8080/tasks")
	if err != nil {
		log.Fatal("Error making HTTP request:", err)
	}
	defer response.Body.Close()

	// Verificar el código de estado de la respuesta
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s\n", response.Status)
	}

	// Leer la respuesta
	var tasks []*Task
	if err := json.NewDecoder(response.Body).Decode(&tasks); err != nil {
		log.Fatal("Error decoding response data:", err)
	}

	// Imprimir todas las tareas
	for _, task := range tasks {
		fmt.Printf("ID: %s, ", task.ID)
		fmt.Printf("Name: %s, ", task.Name)
		fmt.Printf("Status: %s\n", task.Status)
	}
}
