package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage, we need a input value in feet")
		return
	}
	const (
		feetconv = 0.3048
	)

	inputValue := os.Args[1]

	inpflt, err := strconv.ParseFloat(inputValue, 64)

	if err != nil {
		fmt.Println(err)
	}

	valueInMeter := inpflt * 0.3048

	fmt.Printf("The value of %s feets in meters is %f", inputValue, valueInMeter)

}
