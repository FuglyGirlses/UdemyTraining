package main

import "fmt"

func main() {
	a := 10       // declares and assigns (initializes) the value 10 to 'a' of type int.
	b := "golang" // initializes the string "golang" to 'b' of type string.
	c := 4.17     // initializes the value 4.17 to 'c' of type float.
	d := true     // initializes the value 'true' to 'd' of type bool.

	// inside the fmt package use the the Printf function
	// to output the value in variable 'a'.
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	// I'm not completely sure, but it seems %v tells the compiler what kind of data type is being used
	// so that the coder doesn't have to input the data type in the variable declaration.

}
