package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var noKtpPace NoKtp = "123123123123"
	var marriedPace Married = true
	fmt.Println(noKtpPace)
	fmt.Println(marriedPace)
}
