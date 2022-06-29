package main

import "fmt"

func showName(name string) string {
	return "Namaku: " + name
}

func main() {
	total := showName
	fmt.Println(total("Pace"))

}
