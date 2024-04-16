package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//lastMonth := now.AddDate(0, -1, 0)

	date1 := now

	firstOfMonth := time.Date(date1.Year(), date1.Month(), 1, 0, 0, 0, 0, time.Local)
	lastDayOfMonth := firstOfMonth.AddDate(0, 1, -1)

	fmt.Printf("%v %v\n", date1.Month(), date1.Year())

	fmt.Println("Lu Ma Mi Ju Vi Sa Do")
	for i := 0; i < (int(firstOfMonth.Weekday())+6)%7; i++ {
		fmt.Printf("   ")
	}
	for i := 1; i <= lastDayOfMonth.Day(); i++ {
		fmt.Printf("%2d ", i)
		if (i+int(firstOfMonth.Weekday()))%7 == 1 {
			fmt.Println()
		}
	}
	fmt.Println()
}
