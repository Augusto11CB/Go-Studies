package main

import "fmt"

const PI float64 = 3.14

func main() {
	fmt.Println("Hello, World!")

	// --------------------------------

	var name string

	var age int

	var isStudent bool

	fmt.Println(name, age, isStudent)

	// --------------------------------

	var (
		surname          string = "Doe"
		cellphoneNumber  int    = 1234567890
		hasDriverLicense bool   = true
	)

	fmt.Println(surname, cellphoneNumber, hasDriverLicense)

	// --------------------------------

	fmt.Println("Value of PI is", PI)

	// --------------------------------

	var company string = "Wolt"
	fmt.Printf("Value: %s, Type: %T\n", company, company)

	// --------------------------------

	var floatNumber float64 = 10.5
	fmt.Printf("Value: %f, Type: %T\n", floatNumber, floatNumber)

	// --------------------------------

	var onlyPositiveNumber uint32 = 10
	fmt.Printf("Value: %d, Type: %T\n", onlyPositiveNumber, onlyPositiveNumber)

}
