package main

import (
	"fmt"
)

func main() {
	FizBuz(1, 30)
}

// FizBuz is a function that does a Fizz-Buzz
func FizBuz(x, y int) {
	for i := x; i <= y; i++ {
		b := "Buz"
		f := "Fiz"
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(f + b)
		} else if i%3 == 0 {
			fmt.Println(f)
		} else if i%5 == 0 {
			fmt.Println(b)
		} else {
			fmt.Println(i)
		}
	}
}
