package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const maxTurns = 5

func main() {

	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)
		//fmt.Printf("%d", n)

		if n == guess {
			fmt.Println("You win Tan ta da!!!!!!!!!!!!!!!!!!!!!")
			return
		}
	}

	fmt.Println("You Lost, Try again!!!!!!!!!!!!!!!!")
}
