package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var max = os.Args[1]

func main() {

	convMax, err := strconv.Atoi(max)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(convMax)

	fmt.Printf("%5s", "X")

	for i := 0; i >= convMax; i++ {
		fmt.Printf("%/convMax/d", i)
	}
	fmt.Println()

}
