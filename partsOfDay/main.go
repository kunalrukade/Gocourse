package main

import (
	"fmt"
	"time"
)

func main() {

	switch currentHour := time.Now().Hour(); {
	case currentHour >= 20:
		fmt.Println(currentHour)
		fmt.Println("Good Night")
	case currentHour >= 16:
		fmt.Println(currentHour)
		fmt.Println("Have a great Evening")
	case currentHour >= 12:
		fmt.Println(currentHour)
		fmt.Println("Have a lovely Afternoon")
	case currentHour < 12:
		fmt.Println(currentHour)
		fmt.Println("Have a lovely Morning")
	default:
		fmt.Println("There is some error in fetching time")

	}
}
