package main

import (
	"fmt"
	"math/rand"
	//"time"
)

func generateRandomNumber() int {
	//rand.Seed(time.Now().UnixNano())
	return rand.Intn(101)
}

func generateRandomString() string {
	// Caracteres que se usarán para generar el string aleatorio
	caracteres := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Semilla aleatoria basada en la hora actual
	//rand.Seed(time.Now().UnixNano())

	// Longitud aleatoria entre 16 y 32 caracteres
	longitud := rand.Intn(17) + 16

	// Generar el string aleatorio
	b := make([]byte, longitud)
	for i := range b {
		b[i] = caracteres[rand.Intn(len(caracteres))]
	}
	return string(b)
}

func main() {
	fmt.Println("Hello, World!")
	randomNumber := generateRandomNumber()
	fmt.Println("Random Number:", randomNumber)
	randomString := generateRandomString()
	fmt.Println("Random String:", randomString)
}
