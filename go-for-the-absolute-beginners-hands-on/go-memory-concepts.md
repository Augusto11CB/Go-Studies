## Values vs Pointers
### Values (Java doesn't have these)
- Direct data storage (like Java primitives but for ANY type)
- When assigned/passed, the entire data is copied
- Changes to a copy don't affect the original

```
fun main(){

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

	// -------------------------------- Pointers

	person3 := Person{Name: "Bob", Age: 25}

	personPointer := &person3

	personPointer.Name = "Charlie"
	personPointer.Age = 26

	fmt.Println(person3.Name)       // Charlie
	fmt.Println(person3.Age)        // 26
	fmt.Println(personPointer.Name) // Charlie
	fmt.Println(personPointer.Age)  // 26

}
```

### Pointers (Similar to Java references but explicit)
- Store memory addresses instead of actual data
- Use & to get a pointer, * to access the value
- Changes through one pointer affect all references
- *Check `go-pointer-bank-examplo.go` file