package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Println("Cannot convert string!!!!!!!!!!!!!!!", a[1])

	} else {
		fmt.Printf("%d * 2\n", n*2)
	}
}
