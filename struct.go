package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	var pace Customer
	pace.Name = "Ryan Pace"
	pace.Address = "Jl Enggros Kampkey"
	pace.Age = 26

	fmt.Println(pace)

	customer := Customer{
		Name:    "Umar",
		Address: "Depok",
		Age:     26,
	}

	fmt.Println(customer)

}
