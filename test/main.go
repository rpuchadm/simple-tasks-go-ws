// hello world en go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello World")

	longitud := rand.Intn(10)

	// declara un slice de enteros
	numeros := make([]int, longitud)

	// recorre el slice y muestra los elementos
	for i := 0; i < len(numeros); i++ {
		numeros[i] = rand.Intn(100)
	}
	// genera un entero aleatorio entre 0 y 100
	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i])
	}

	fmt.Println("-------------")

	// declara la fecha de hoy
	today := time.Now()
	nextweek := time.Now().Add(time.Hour * 24 * 5)

	fmt.Println("Hoy es: ", today)
	fmt.Println("La fecha de la próxima semana es: ", nextweek)

	days := calculateDaysBetweenDates(today, nextweek)
	fmt.Println("Días entre hoy y la próxima semana: ", days)

	fmt.Println("-------------")

	// declara un mapa de strings
	textos := make(map[string]string)

	// agrega elementos al mapa
	textos["uno"] = "Hola"
	textos["dos"] = "Mundo"

	// recorre el mapa y muestra los elementos
	for key, value := range textos {
		fmt.Println(key, value)
	}

}

func calculateDaysBetweenDates(today time.Time, nextweek time.Time) int {
	// calcula la diferencia de días entre dos fechas
	days := nextweek.Sub(today).Hours() / 24
	return int(days)
}
