package main

import "fmt"

func main() {
	fmt.Println("Data Structure - Arrays")

	var employees = [3]string{"John", "Jane", "Jim"}
	var employeesAge = [3]int{25, 26, 27}

	fmt.Println(employees, employeesAge)
	fmt.Println(len(employees), len(employeesAge))
	fmt.Printf("Type of employees: %T, Type of employeesAge: %T\n", employees, employeesAge)

	// replacing the value of the array
	fmt.Println("Replacing the value of the array")
	employees[0] = "Peter"
	employeesAge[0] = 33

	fmt.Println(employees, employeesAge)

	// declaring an array with no values
	var workers = [3]string{}
	fmt.Println(workers)

	// -------------------------------- Arrays - Initialization with index and value
	var fruits = [3]string{0: "Apple", 2: "Cherry"}
	fmt.Println(fruits)

	// -------------------------------- Values vs Pointers - Value

	type Person struct {
		Name string
		Age  int
	}

	person1 := Person{Name: "Alice", Age: 30}

	// Complete copy of person1 is created
	person2 := person1

	// Modifies only person2, person1 is unchanged
	person2.Name = "Augusto"
	person2.Age = 150

	fmt.Println(person1.Name) // Alice
	fmt.Println(person1.Age)  // 30
	fmt.Println(person2.Name) // Augusto
	fmt.Println(person2.Age)  // 31

	// -------------------------------- Values vs Pointers - Pointer

	person3 := Person{Name: "Bob", Age: 25}

	personPointer := &person3

	personPointer.Name = "Charlie"
	personPointer.Age = 26

	fmt.Println(person3.Name)       // Charlie
	fmt.Println(person3.Age)        // 26
	fmt.Println(personPointer.Name) // Charlie
	fmt.Println(personPointer.Age)  // 26

	// --------------------------------

}
