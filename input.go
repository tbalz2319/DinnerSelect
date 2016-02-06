package main

import "fmt"

func main() {
	var myname string
	str1 := "What is your name?"
	fmt.Println(str1)
	fmt.Scanf("%s", &myname)
	fmt.Println("Hello", myname)
	str2 := "Do you like cookies?"
	fmt.Println(str2, myname)
}
