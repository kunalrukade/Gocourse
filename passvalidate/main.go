package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	if len(args) != 3 {
		fmt.Println("Usage:[username] [password]")
		return
	}

	userName := os.Args[1]
	password := os.Args[2]

	if userName != "Kunal" && userName != "Pooja" {
		fmt.Printf("Access denied for %q.\n", userName)
	} else if userName == "Kunal" && password == "1234" || userName == "Pooja" && password == "5678" {
		fmt.Printf("Access granted to %q.\n", userName)
	} else {
		fmt.Println("Error with password")
	}

}
