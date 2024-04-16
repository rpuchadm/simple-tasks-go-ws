package main

import (
	"fmt"
	"math/rand"
)

func generateRandomString() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {

	randomNumber := rand.Intn(10) + 1
	fmt.Println("Número aleatorio:", randomNumber)

	if randomNumber > 5 {

		stringsArray := make([]string, randomNumber)
		for i := 0; i < randomNumber; i++ {
			stringsArray[i] = generateRandomString()
		}
		fmt.Println("Strings Array:", stringsArray)
		fmt.Println("Strings Array:")
		fmt.Println("<ul>")
		for _, str := range stringsArray {
			fmt.Println("<li>", str, "</li>")
		}
		fmt.Println("</ul>")

	} else {

		array := make([]int, randomNumber)
		for i := 0; i < randomNumber; i++ {
			array[i] = rand.Intn(100)
		}
		fmt.Println("Array:", array)
		fmt.Println("Array:")
		fmt.Println("<ul>")
		for _, num := range array {
			fmt.Println("<li>", num, "</li>")
		}
		fmt.Println("</ul>")

	}

}
