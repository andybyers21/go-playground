/* If we list all the natural numbers below 10 that are multiples of 3 or 5, we
   get 3, 5, 6 and 9. The sum of these multiples is 23.

   Find the sum of all the multiples of 3 or 5 below 1000. */

package main

import "fmt"

func prob(many int) {
	var nums []int
	var mults []int
	var result int

	// append numbers divisable by 3 or 5 to a slice
	for i := 0; i < many; i++ {
		nums = append(nums, i)
		if i%3 == 0 || i%5 == 0 {
			mults = append(mults, i)
		}
	}

	// add all numbers divisable by 3 or 5 together
	for n := 0; n < len(mults); n++ {
		result = result + mults[n]
	}

	fmt.Println(result)

}

func main() {
	prob(1000)
}
