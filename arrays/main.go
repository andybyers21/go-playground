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
	var score [5]float64
	score[0] = 30
	score[1] = 14
	score[2] = 22
	score[3] = 28
	score[4] = 4

	var percentScore [5]float64
	var t float64

	for i := 0; i < 5; i++ {
		percentScore[i] = (score[i] / 30) * 100
		percentScore[i] = math.Round(percentScore[i])
		t += percentScore[i]
	}

	averageScore := t / len(score)

	fmt.Println("all test scores: ", score)
	fmt.Println("all scores by %: ", percentScore)
	fmt.Println("average test score: ", averageScore, "%")
}
