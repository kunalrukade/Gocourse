package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type user struct {
	Name        string          `json:"userName"`
	Permissions map[string]bool `json:"Permission,omitempty"`
}

func main() {

	var input []byte

	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
		input = append(input, '\n')
	}

	var users []user

	if err := json.Unmarshal(input, &users); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(users)

	for _, users := range users {
		fmt.Print("+", users.Name)

		switch p := users.Permissions; {
		case p == nil:
			fmt.Print("has no power")
		case p["admin"]:
			fmt.Print("is a admin")
		case p["write"]:
			fmt.Print("can write")
		}
	}

}
