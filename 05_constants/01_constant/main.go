package main

import "fmt"

const x string = "I am a constant"

func main() {
	const y = 42
	fmt.Println("x - ", x) // outputs x - i am a constant
	fmt.Println("y - ", y) // outputs y - 42
}
