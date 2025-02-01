package main

import (
	"fmt"
	"time"
)

func main() {
	dateStr := "20-02-2025"

	taskDate, _ := time.Parse("01-02-2006", dateStr)
	fmt.Println("create task date input after format taskDate", taskDate)

	dateStr2 := taskDate.Format("02-01-2006")

	fmt.Println("create task date input after format", dateStr2)
}
