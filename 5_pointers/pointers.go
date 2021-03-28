package main

import (
	"fmt"
)

func main() {
	basicSyntax()
}

func basicSyntax() {
	// Decalre some variables
	x, y := 10, 10000
	// Print the variables
	fmt.Println("Variables x and y: ", x, y)
	// Print the ADDRESS OF the variables using `&`
	fmt.Println("The address of variables x and y: ", &x, &y)

	// Create a pointer: Assign the address of (`&`) to a new variable
	// The value of `a` here is the address of `x`
	a := &x
	// Println a will return the address of `x`
	fmt.Println("Variable a has been assigned the address of x (&x): ", a)
	// A `*` acts as an operator and returns the value of whateveer the pointer
	// variable is pointing to.
	fmt.Println("By printing the pointer, *a, you will return the value of x (&x): ", *a)
	// check this by printing the type of `a`
	fmt.Printf("Which is an: %T\n", a)

	// You can then the value of the pointer which will change the value
	// of the original variable.
	b := &y
	*b = 555
	fmt.Println("Modify the value of variable y by changing the value of the pointer *b (`*b = 555`): ", y)
	// If we do something to `*b`, we are modifying the value of `y`
	*b += 1
	fmt.Println("You can then carry out operations on the variable by modifying the pointer (`*b += 1`): ", y)
}
