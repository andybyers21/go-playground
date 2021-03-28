package main

import (
	"fmt"
	"math"
)

// The `type` command creates a new type. We are defining the struct of
// "Circle"
type Circle struct {
	x, y, r float64
}

func main() {
	// create a new instance of `Circle`, a local circle. By default the values
	// will be set to 0
	roundThing := Circle{
		x: 4,
		y: 2,
		r: 3,
	}

	// Print the area of `roundThing` to 2 decimal places using the circleArea function.
	// (Print the address of `roundThing` as a pointer has been used to modify
	// the original value in `circleArea`)
	fmt.Printf("%.2f\n", circleArea(&roundThing))

}

func circleArea(cr *Circle) float64 {
	// Inputs a `Circle`, returns area as a float64
	// Use a pointer to `Circle` to modify the original values (in `main()` in
	// this case).
	// Takes the radius of a defined `Circle`, calculates and returns the area
	area := math.Pi * (cr.r * cr.r)
	return area
}
