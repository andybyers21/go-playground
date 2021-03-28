package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	scoreCheck()
	fmt.Println("all test scores: ", score)
	percentageConvert()
	smallLarge()
}

// Test Scores
// ---- ------
var score = []float64{
	30, 14, 22, 28, 4, 12, 26, 29, 18,
	26, 13, 20, 4, 21, 29, 22, 15, 7,
	26, 12, 17, 15, 19, 21, 26, 28, 29,
}

// Check initial scores for errors
func scoreCheck() {
	// Loop through code and find scores greater than 30 or less than 0. If
	// found exit with status code 1.
	for i := 0; i < len(score); i++ {
		if score[i] > 30 || score[i] < 0 {
			fmt.Println("Incorrect score submitted (out of range):", score[i])
			os.Exit(1)
		}
	}
}

// Convert test scores out of 30 (in a slice) to a percentage:
func percentageConvert() {
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
	fmt.Println("all scores by %: ", percentScore)
	fmt.Println("average test score: ", averageScore, "%")
}

// find the smallest and largest test scores.
func smallLarge() {
	small := score[0]
	large := score[0]
	for i := 0; i <= len(score)-1; i++ {
		if score[i] < small {
			small = score[i]
		} else if score[i] > large {
			large = score[i]
		}
	}

	smallPercent := math.Round((small / 30) * 100)
	largePercent := math.Round((large / 30) * 100)

	fmt.Println("lowest score was: ", small, "/", smallPercent, "%", "\nhighest score was: ", large, "/", largePercent, "%")
}
