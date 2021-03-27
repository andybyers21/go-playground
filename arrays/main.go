// Arrays, Slices and Maps

package main

import (
	"fmt"
	"math"
)

func main() {
	percentageConvert()
}

func percentageConvert() {
	// convert test scores out of 30 (in an array) to a percentage

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
