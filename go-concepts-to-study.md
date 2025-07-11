## Core Go Interface Concepts
### 1. Structural Typing vs Nominal Typing
Structural typing (Go) - types are compatible based on their structure/shape
Nominal typing (Java/Kotlin) - types must be explicitly declared as related
Search: "Go structural typing", "structural vs nominal typing"

### 2. Implicit Interface Satisfaction
Go's key feature: types automatically satisfy interfaces without explicit declaration
Search: "Go implicit interface satisfaction", "Go automatic interface implementation"

### 3. Duck Typing
"If it walks like a duck and quacks like a duck, it's a duck"
Go implements compile-time duck typing through structural typing
Search: "Go duck typing", "compile-time duck typing"

## Specific Go Interface Features
### 4. Interface Embedding

```
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type ReadWriter interface {
    Reader  // Embedded interface
    Writer  // Embedded interface
}
```
Search: "Go interface embedding", "Go interface composition"

### 5. Empty Interface
```
interface{} // Can hold any type
any         // Go 1.18+ alias for interface{}
```
Search: "Go empty interface", "Go interface{}", "Go any type"

### 6. Type Assertions and Type Switches

```
// Type assertion
if v, ok := x.(string); ok {
    // x is a string
}

// Type switch
switch v := x.(type) {
case string:
    // v is string
case int:
    // v is int
}
```
Search: "Go type assertions", "Go type switches"

## Advanced Interface Patterns
### 7. Interface Segregation
Multiple small, focused interfaces vs large monolithic ones
Search: "Go interface segregation", "Go small interfaces"

### 8. Interface Composition Patterns
Building complex interfaces from simple ones
Search: "Go interface composition patterns"

### 9. Polymorphism in Go
How Go achieves polymorphism without traditional inheritance
Search: "Go polymorphism", "Go polymorphism without inheritance"


### Specific Web Searches:
```
"Go interfaces structural typing"
"Go implicit interface implementation"
"Go interface satisfaction rules"
"Go duck typing compile time"
"Go interface embedding composition"
"Go polymorphism interfaces"
"Go type system vs Java type system"
"Go interfaces vs Java interfaces"
```

### Key Concepts to Understand:
Interface Values: How Go stores both type and value information
Method Sets: Rules about which methods belong to which types
Pointer vs Value Receivers: How they affect interface satisfaction
Interface Nil Values: Understanding nil interfaces vs nil concrete values

### Comparative Studies:
Search for articles comparing:
```
"Go interfaces vs Java interfaces"
"Go vs Java polymorphism"
"Structural typing vs nominal typing examples"
"Go type system design philosophy"
```
