package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	kunal := person{
		firstName: "Kunal",
		lastName:  "Avghade",
		age:       25,
	}
	// or
	kunal2 := person{"Kunal", "Avghade", 25}

	fmt.Printf("%+v\n", kunal)
	fmt.Printf("%+v\n", kunal2)
}
