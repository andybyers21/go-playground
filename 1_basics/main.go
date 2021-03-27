package main

import (
	"fmt"
)

func main() {
	upToNum(15)
}

func upToNum(x int) {
	for i := 1; i <= x; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

func farToC() {
	fmt.Println("Please enter farinheight to be converted: ")
	var f float64
	fmt.Scanf("%f", &f)
	var c float64 = (f - 32) * 5 / 9
	fmt.Printf("%.2f", c)
}

func feetToMeters() {
	fmt.Println("enter feet to be converted: ")
	var f float64
	fmt.Scanf("%f", &f)
	var m float64 = f * 0.3048
	fmt.Printf("%.2f", m)
}
