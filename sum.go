package main

import "fmt"

//global variable declartation
var a int = 20

func main() {
	//local variable declaration
	var a int = 100
	var b int = 20
	var c int = 0
	var m int = 0
	var d int = 0
	var s int = 0
	var h int = 0

	fmt.Printf("value of a in main() = %d\n", a)
	c = sum(a, b)
	m = mult(a, b)
	d = div(a, b)
	s = sub(a, b)
	h = hundredplus(a, b)
	fmt.Printf("value of c in main() = %d\n", c)
	fmt.Printf("The value of m in main() = %d\n", m)
	fmt.Printf("The value of d in main() = %d\n", d)
	fmt.Printf("The two numbers subracted equal = %d\n", s)
	fmt.Printf("The two numbers added + 100 = %d\n", h)
}

//The function below adds the variables below based on locally
//declared variables from the main section of the program
func sum(a, b int) int {
	fmt.Printf("Value of a in sum() = %d\n", a)
	fmt.Printf("value of b in sum() = %d\n", b)

	return a + b
}

//The function below multiplies two variables and returns the results
func mult(a, b int) int {
	fmt.Printf("Value of a in mult = %d\n", a)
	fmt.Printf("value of b in mult() = %d\n", b)

	return a * b
}

//The function below multiplies two variables and returns the results
func div(a, b int) int {
	return a / b
}

func sub(a, b int) int {
	return a - b
}

func hundredplus(a, b int) int {
	return (a + b) + 100

}
