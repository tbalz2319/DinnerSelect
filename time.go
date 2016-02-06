package main

import "fmt"
import "time"

func main() {
	fmt.Println(time.Now().Format(time.RFC850))
	current_time := time.Now().Local()
	fmt.Println("The Current time is ", current_time.Format("20160201"))
}
