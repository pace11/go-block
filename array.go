package main

import "fmt"

func main() {
	var name [2]string
	name[0] = "Ryan"
	name[1] = "Pace"
	fmt.Println(name[0])
	fmt.Println(name[1])

	var values = [3]int{
		100,
		50,
		45,
	}
	fmt.Println(values)
}
