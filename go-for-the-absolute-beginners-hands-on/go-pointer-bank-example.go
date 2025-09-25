// package gofortheabsolutebeginnershandson
package main

// We’ll model a bank account system with:
// 	•	a struct to hold account data,
// 	•	methods that update the struct (using pointer receivers),
// 	•	an interface for different account types,
// 	•	and a function that works with the interface.

import "fmt"

// Account is an interface for bank accounts
type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	GetBalance() float64
}

// BankAccount is a struct that implements Account
type BankAccount struct {
	Owner   string
	Balance float64
}

// Deposit increases balance (pointer receiver needed to modify the struct)
func (a *BankAccount) Deposit(amount float64) {
	a.Balance += amount
}

// Withdraw decreases balance if enough funds are available
func (a *BankAccount) Withdraw(amount float64) {
	if a.Balance >= amount {
		a.Balance -= amount
	} else {
		fmt.Println("Insufficient funds")
	}
}

// GetBalance returns the balance (no modification, so value receiver is fine)
func (a BankAccount) GetBalance() float64 {
	return a.Balance
}

// A function that works with any Account
func PrintAccountInfo(acc Account) {
	fmt.Printf("Balance: %.2f\n", acc.GetBalance())
}

func main() {
	// Create a new account
	acc := BankAccount{Owner: "Alice", Balance: 100}

	fmt.Println("Initial balance in struct:", acc.Balance) // 100

	// Use interface functions
	var a Account = &acc // must pass pointer to satisfy Deposit/Withdraw methods

	a.Deposit(50) // modifies balance
	a.Withdraw(30)

	fmt.Println("Final balance in struct:", acc.GetBalance()) // 120
}

// Why pointers are important here:
// 	1.	Deposit and Withdraw use pointer receivers (*BankAccount) so they can modify the actual balance inside the struct.
// 	•	If we used value receivers, Go would copy the struct, and changes would be lost.
// 	2.	In main, when we assign to the interface var a Account = &acc, we must pass a pointer, since only the pointer type implements those methods.
// 	3.	For GetBalance, a value receiver is fine since it doesn’t modify the struct.

// This is the classic Go pattern: use pointers when the method should modify state, use values for read-only methods.

// Want me to also extend this with another struct (like SavingsAccount) so you see how multiple types can satisfy the same interface?
