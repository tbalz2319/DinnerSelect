package main

import "fmt"

func main() {
	// Local variable declaration
	var a int = 7568
	var b int = 2000
	var ret int

	// Calling a function to get the max value
	ret = max(a, b)

	fmt.Printf("The Max Value is : %d\n", ret)
}

// The max function that returns the max of two numbers is below
func max(num1, num2 int) int {
	// Local variable declaration for max function below
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
