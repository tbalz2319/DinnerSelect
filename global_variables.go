package main

import "fmt"

//global variable declaration
var g int

func main() {

	//local variable declaration
	var a, b int

	//actual variable initialization
	a = 100
	b = 206
	g = a + b

	fmt.Printf("The value of a = %d, b = %d and g = %d\n", a, b, g)
}
