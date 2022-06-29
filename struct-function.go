package main

import "fmt"

type Human struct {
	Name, Address string
	Age           int
}

func (human Human) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", human.Name)
}

func (a Human) sayNo() {
	fmt.Println("Sad", a.Name)
}

func main() {

	var pace Human
	pace.Name = "Umary"

	pace.sayHi("Pace")
	pace.sayNo()

}
