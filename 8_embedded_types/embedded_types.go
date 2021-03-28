package main

import "fmt"

// 1 create a struct to hold "person" details
type person struct {
	fName string
	lName string
	age   int
}

// 2 crete a struct that holds secret agent fields and embedds person type
type secretAgent struct {
	person
	license bool
}

// 3 attach a method to person (pSpeak)
func (p person) speak() {
	fmt.Println("uptown funky-wunk")
}

// 4 attach a method to secret agent (saSpeak)
func (sa secretAgent) speak() {
	fmt.Println("Shaken not stirred")
}

func main() {
	// 5 create a variable of type person and print a field
	person1 := person{
		fName: "Bill",
		lName: "Billson",
		age:   79,
	}
	fmt.Println(person1)

	// 6 create a variable of type secret agent and print a field
	// Notice that type person is embedded into secretAgent in this way and
	// person must be specified as... :
	agent1 := secretAgent{
		person: person{
			fName: "James",
			lName: "Pond",
			age:   7,
		},
		license: true,
	}
	fmt.Println(agent1)

	// 7 run speak attachd to the variable type person
	person1.speak()

	// 8 run speak attachd to the variable type secret agent
	agent1.speak()

	// 9 run speak as person attached to secret agent
	agent1.person.speak()
}
