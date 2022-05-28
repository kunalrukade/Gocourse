package main

import "fmt"

type Book struct {
	title string
	price float64
}

func PrintBook(b Book) {
	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
}
