Go doesn't have traditional object-oriented concepts like classes and inheritance, but it does have features that support object-oriented programming patterns:
1. Instead of classes, Go uses structs to define custom data types
2. Methods can be associated with any type through receiver functions
3. Go uses composition instead of inheritance
4. Interfaces provide polymorphism through implicit implementation
5. Factory functions serve as constructors

```
// Define a struct (instead of class) (1)
type Person struct {
    Name string
    Age  int
}

// Associate methods with types (2)
func (p Person) Greet() string {
    return "Hello, I'm " + p.Name
}


// "Instantiate" a struct (5)
func NewPerson(name string, age int) Person {
    return Person{Name: name, Age: age}
}
```

```

// Composition instead of inheritance (3)
// Base "class" equivalent
type Vehicle struct {
    MaxSpeed int
    Weight   float64
}

func (v Vehicle) GetMaxSpeed() int {
    return v.MaxSpeed
}

type Car struct {
    Vehicle    // Embedding Vehicle (composition)
    Doors      int
    EngineType string
}

// Usage
func main() {
    myCar := Car{
        Vehicle:    Vehicle{MaxSpeed: 180, Weight: 1500.5},
        Doors:      4,
        EngineType: "V6",
    }
    
    // Car "inherits" Vehicle methods
    fmt.Println("Max speed:", myCar.GetMaxSpeed())
}
```

```
// Interfaces for polymorphism (4)

// Define interface
type Speaker interface {
    Speak() string
}

// Types implicitly implement the interface
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return d.Name + " says woof!"
}

type Cat struct {
    Name string
}

func (c Cat) Speak() string {
    return c.Name + " says meow!"
}

// Polymorphic function
func MakeSpeak(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    dog := Dog{Name: "Rex"}
    cat := Cat{Name: "Whiskers"}
    
    // Both types can be used wherever Speaker interface is accepted
    MakeSpeak(dog)  // Rex says woof!
    MakeSpeak(cat)  // Whiskers says meow!
}
```