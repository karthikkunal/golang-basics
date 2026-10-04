// In Go, a struct is a composite data type that groups together zero or
// more fields with different data types under a single name.
// It's similar to a class in other languages, but without methods.
// Each field in a struct is accessed using dot notation.

// Here's a simple example to illustrate how structs work in Go:

package main

import "fmt"

// Define a struct named 'Person' with two fields: 'name' and 'age'
type Person struct {
	name string
	age  int
}

func main() {
	var p Person // This declares a variable named p of type Person.
	// This variable will hold information about a person,
	// including their name and age.
	p.name = "William"         // Assign values
	p.age = 20                 // Assign values
	fmt.Println(p.name, p.age) // this is to print the output
}

// Output:
// William 20
