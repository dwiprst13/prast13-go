package main

import "fmt"

func main() {
	name := "Dwi Prasetia"
	var num1 int = 12
	var num2 int = 245
	var numSum = num1 + num2
	var numMulti = num1 * num2
	var numSubs = num1 - num2
	var status = !false

	fmt.Println("value from ", num1, "+", num2, "=", numSum)
	fmt.Println("value from ", num1, "x", num2, "=", numMulti)
	fmt.Println("value from ", num1, "-", num2, "=", numSubs)
	fmt.Println(status)
	fmt.Println(name)
}
