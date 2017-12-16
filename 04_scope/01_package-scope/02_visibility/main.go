package main

import "fmt"

var x string = "I am visible to all" // package variable - global

func main() {
	a := 42                                 // block variable - visible only to func main()
	b := "I am only visible in func main()" // block variable - visible only to func main()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(x)

	vis()
}

func vis() {
	c := 10                                // block variable - visible only to func vis()
	d := "I am only visible in func vis()" // block variable - visible only to func vis()

	fmt.Println(c)
	fmt.Println(d)
	// fmt.Println(b)

}
