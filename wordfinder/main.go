package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "" + "lazy car jumps again and again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

	for _, q := range query {
		for i, w := range words {
			if strings.ToUpper(q) == strings.ToUpper(w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				break
			}
		}
	}
	fmt.Println("This word isnt found in the Corpus")
}
