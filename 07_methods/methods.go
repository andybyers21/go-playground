package main

import (
	"fmt"
	"math"
)

// Define Circle struct
type Circle struct {
	x, y, r float64
}

/* -------------------------------------------------------------------------------
Convert the function to a method:
	Add the receiver(s) (between the func keyword and the name).
	The receiver is like a parameter but by writing it this way it allows us to
	use .dot-notation
------------------------------------------------------------------------------- */
func (cr *Circle) circleArea() float64 {
	// Inputs a `Circle`, returns area as a float64
	// Use a pointer to `Circle` to modify the original values (in `main()` in
	// this case).
	// Takes the radius of a defined `Circle`, calculates and returns the area
	return math.Pi * (cr.r * cr.r)
	// have also refactored this into a one liner
}

// Now lets make a method to calculate the area of a rectangle (A = w*l)
// First construct the struct
type Rectangle struct {
	w, l float64
}

// then create the method
// `ra` is whatever we difine as w and l, pointer to circle
func (ra *Rectangle) rectangleArea() float64 {
	return ra.w * ra.l
}

func main() {
	// create a new instance of `Circle`, a local circle.
	// lets tidy this up a bit so we can calculate and print in 2 lines.
	roundThing := Circle{x: 4, y: 2, r: 3}

	// Print the area of `roundThing` to 2 decimal places using the circleArea function.
	// (Print the address of `roundThing` as a pointer has been used to modify
	// the original value in `circleArea`)
	// fmt.Printf("%.2f\n", circleArea(&roundThing))

	// Once we create a method we can print the result with dot-notation
	// Easier to read and we no longer need the `&` operator as go automatically
	// knows to to pass a pointer to the `rountThing`.
	fmt.Printf("Circle area = %.2f\n", roundThing.circleArea())

	// Rectangle
	squareish := Rectangle{w: 10, l: 15}
	fmt.Printf("Rectangle area = %.2f\n", squareish.rectangleArea())
}
