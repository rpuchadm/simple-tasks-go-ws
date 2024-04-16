// main.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Verificar si se proporciona un ID de tarea como argumento
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <task_id>")
		os.Exit(1)
	}

	taskID := os.Args[1]

	// Realizar la solicitud HTTP para eliminar la tarea
	req, err := http.NewRequest("DELETE", "http://localhost:8080/tasks/"+taskID, nil)
	if err != nil {
		log.Fatal("Error creating HTTP request:", err)
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making HTTP request:", err)
	}
	defer response.Body.Close()

	// Verificar el código de estado de la respuesta
	if response.StatusCode != http.StatusNoContent {
		log.Fatalf("Error: %s\n", response.Status)
	}

	fmt.Println("Task deleted successfully")
}
