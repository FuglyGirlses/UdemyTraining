package main

import "fmt"

// var 'x' is accessible to any function in this package.

var x int = 42 // initialization of x.

// main() function
func main() {
	fmt.Println(x)   // print the value of x
	fmt.Println("x") // prints the letter 'x'
	foo()            // call the foo function
}

// foo() function
func foo() {
	fmt.Println(x) // print the value of x
}
