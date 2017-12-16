package main

import "fmt" // fmt is not package visible and must be imported for every file.

// var x is visible to all
var x string = "outside func main()"

func main() {
	fmt.Println(x)
	y := "inside func main()"
	fmt.Println(y)
	{ // var z is only visible between these braces
		z := "inside, inside func main()"
		fmt.Println(z) // only visible here
		fmt.Println(y) // y is initialized outside so its visible inside the braces
		fmt.Println(x) // x is initialized outside so its visible inside the braces
	}
	//fmt.Println(z) // z is initialized inside the braces so it is not visible outside them
}
