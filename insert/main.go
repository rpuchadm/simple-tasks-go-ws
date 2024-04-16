// main.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Verificar si se proporciona un nombre de tarea como argumento
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go \"Task Name\"")
		os.Exit(1)
	}

	taskName := os.Args[1]

	// Crear la estructura de datos para la tarea
	task := struct {
		Name string `json:"name"`
	}{
		Name: taskName,
	}

	// Convertir la estructura de datos en JSON
	requestBody, err := json.Marshal(task)
	if err != nil {
		log.Fatal("Error marshaling task data:", err)
	}

	// Realizar la solicitud HTTP para agregar la tarea
	response, err := http.Post("http://localhost:8080/tasks", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal("Error making HTTP request:", err)
	}
	defer response.Body.Close()

	// Leer la respuesta
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response data:", err)
	}

	// Verificar el código de estado de la respuesta
	if response.StatusCode != http.StatusCreated {
		log.Fatalf("Error: %s\n", responseData)
	}

	// Extraer el ID de la tarea de la respuesta
	var taskResponse struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(responseData, &taskResponse); err != nil {
		log.Fatal("Error unmarshaling response data:", err)
	}

	// Imprimir el ID de la tarea
	fmt.Println("Task created with ID:", taskResponse.ID)
}
