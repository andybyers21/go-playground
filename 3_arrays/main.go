// Compute test scores using Arrays & Slices

package main

import (
	"fmt"
	"math"
)

func main() {
	// percentageConvert()
	smallestInt()
}

// Write a program that finds the smallest number in an array:
func smallestInt() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	// Assign the first number in arrax x to variable h. Loop through each
	// number in the array and if smaller than h replace h. Once the loop is
	// complete this will leave the smallest number in var h
	h := x[0]
	for i := 0; i <= len(x)-1; i++ {
		if x[i] < h {
			h = x[i]
		}
	}

	fmt.Println(h)
}

// convert test scores out of 30 (in an array) to a percentage
func percentageConvert() {

	// test scores
	score := []float64{
		30, 14, 22, 28, 4, 12, 26, 29, 18,
	}

	// create a slice to hold percentage scores
	var percentScore = make([]float64, len(score))

	// initialize a variable to calculate average score
	var t float64

	for i := 0; i < len(score); i++ {
		// calculate scores as a percentage
		percentScore[i] = (score[i] / 30) * 100
		// round scores to nearest int
		percentScore[i] = math.Round(percentScore[i])
		// add all scores together
		t += percentScore[i]
	}
	// calculate average of all scores
	averageScore := t / float64(len(score))
	// round to a whole number
	averageScore = math.Round(averageScore)

	// print results to the terminal
	fmt.Println("all test scores: ", score)
	fmt.Println("all scores by %: ", percentScore)
	fmt.Println("average test score: ", averageScore, "%")
}
