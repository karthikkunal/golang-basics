// In Go, an interface is a type that specifies a set of method signatures.
// Any type that implements all the methods of an interface implicitly implements that interface.
// This allows for polymorphism and decoupling of code.

// Here's a simple example demonstrating the use of interfaces in Go

package main

import "fmt"

// Define an interface named Animal
type Animal interface {
	Sound() string
}

// Define a struct named Dog
type Dog struct{}

// Implement the Sound method for Dog
func (d Dog) Sound() string {
	return "Woof"
}

// Define a struct named Cat
type Cat struct{}

// Implement the Sound method for Cat
func (c Cat) Sound() string {
	return "Meow"
}

func main() {
	// Declare a variable of type Animal
	var animal Animal

	// Create a Dog instance and assign it to the animal variable
	animal = Dog{}
	fmt.Println("Dog says:", animal.Sound())

	// Create a Cat instance and assign it to the animal variable
	animal = Cat{}
	fmt.Println("Cat says:", animal.Sound())
}

// go run main.go
// Output:
// Dog says: Woof
// Cat says: Meow
