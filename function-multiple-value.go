package main

import "fmt"

func sayHello() (string, string, string) {
	return "Ryan", "Putra", "Pace"
}

func main() {
	fmt.Println(sayHello())
}
