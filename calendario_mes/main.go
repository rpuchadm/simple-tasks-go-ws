package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	lastDayOfMonth := firstOfMonth.AddDate(0, 1, -1)

	fmt.Printf("%v %v\n", now.Month(), now.Year())

	fmt.Println("Lu Ma Mi Ju Vi Sa Do")
	// Ajustamos el cálculo para el primer día de la semana
	for i := 0; i < (int(firstOfMonth.Weekday())+6)%7; i++ {
		fmt.Printf("   ")
	}
	for i := 1; i <= lastDayOfMonth.Day(); i++ {
		fmt.Printf("%2d ", i)
		if (i+int(firstOfMonth.Weekday()))%7 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
