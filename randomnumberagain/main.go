package main

func main() {
	guess := 10

	for n := 0; n != guess {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d", n)
	}
}