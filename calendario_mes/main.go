// hello world
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	lastDayOfMonth := time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, time.Local)

	fmt.Printf("%v %v\n", now.Month(), now.Year())

	fmt.Println("Lu Ma Mi Ju Vi Sa Do")
	for i := 0; i < int(now.Weekday()); i++ {
		fmt.Printf("   ")
	}
	for i := 1; i <= lastDayOfMonth.Day(); i++ {
		fmt.Printf("%2d ", i)
		if (i+int(now.Weekday()))%7 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
