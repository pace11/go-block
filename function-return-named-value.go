package main

import "fmt"

func sayHello() (firstName, lastName string) {
	firstName = "Ryan"
	lastName = "Pace"
	return
}

func main() {
	fmt.Println(sayHello())
}
