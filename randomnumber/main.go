package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5
	usage    = `Welcome to the Luck number Game!!!!!!!
	
	The Program will generate %d random number, your mission is guess one of those number
	
	The Greater your number is, the narder it gets.
	
	Wanna Play?`
)

func main() {

	rand.Seed(time.Now().UnixNano())
	//guess := 25

	args := os.Args[1:]
	if len(os.Args) != 2 {
		fmt.Printf(usage, maxTurns)
		return

	}

	guess, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatal(err)
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive integer")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("You win, Congrats you get a car!!!!!!!!!!!")
			return
		}

	}
	fmt.Println("You loose....... Try Again")
}
