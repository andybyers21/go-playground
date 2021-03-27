// scoreays, Slices and Maps

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
		30, 14, 22, 28, 4,
	}

	var percentScore []float64
	var t float64

	for i := 0; i < len(score); i++ {
		percentScore[i] = (score[i] / 30) * 100
		percentScore[i] = math.Round(percentScore[i])
		t += percentScore[i]
	}

	averageScore := t / float64(len(score))

	fmt.Println("all test scores: ", score)
	fmt.Println("all scores by %: ", percentScore)
	fmt.Println("average test score: ", averageScore, "%")
}
