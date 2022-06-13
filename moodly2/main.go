package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	name := args[0]
	feelings := args[1]

	moods := [2][3]string{
		{"Happy", "Awesome", "Amazed"},
		{"Sad", "Aweful", "Lonely"},
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods))

	if feelings == "positive" {
		fmt.Printf("%s is %s\n", name, moods[0][n])
	} else {
		fmt.Printf("%s is %s", name, moods[1][n])
	}

}
