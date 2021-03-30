package main

import (
	"fmt"
	"testing"
)

// Define the common behaviour between structs
type voice interface {
	sound()
}

// Define structs that make a sound()
type bird struct {
	species string
	color   string
}

type car struct {
	model string
	color string
}

type human struct {
	name   string
	gender string
}

// Define soudnds that structs make
func (b bird) sound() {
	fmt.Println("tweet, tweet")
}

func (c car) sound() {
	fmt.Println("brum, brum")
}

func (h human) sound() {
	fmt.Println("Hello there, I am a human")
}

func testBird(t *testing.T) {
	t.Error("expected tweet, tweet")
}

// ...
func main() {
	bird1 := bird{"starling", "brown"}
	car1 := car{"escort", "white"}
	human1 := human{"dave", "female"}

	bird1.sound()
	car1.sound()
	human1.sound()
}
