package main

import "fmt"

func main() {
	// multiple initialization
	const (
		p = 3.14
		q = "i am a string"
	)
	fmt.Println(p, " - ", q) // outputs 3.14 - i am a string
}
