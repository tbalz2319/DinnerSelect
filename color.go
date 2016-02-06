package main

import "fmt"

func main() {
	fmt.Println("What is your favorite color?")
	var favColor string
	fmt.Scanf("%s", &favColor)
	fmt.Println("Your favorite color is", favColor)

}
