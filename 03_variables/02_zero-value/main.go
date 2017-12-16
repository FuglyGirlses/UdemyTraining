package main

import "fmt"

func main() {
	// setting variables to a zero-value
	var a int /// tells the compiler that type int data will be stored in 'a'.
	var b string
	var c float64
	var d bool

	// outputs the values stored in the variables a,b,c,d which is 0.
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}
