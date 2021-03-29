/* Normally when we run a function the program will scan a line, wait for it to
complete then move onto the next one.

Using a GO routine it will not wait and immediately run the next opersation */

package main

import "fmt"

func f(n int) {
	// Create a for loop that runs some arbitary code
	// In this case format a string with teh given numbers
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	// Using `go f()` will run the calculations concurrently.
	go f(0)
	var input string
	fmt.Scanln(&input)
	// Note that this program will not stop running until it hears some kind of
	// user input.
}
