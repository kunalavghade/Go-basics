package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	age       int
}

func main() {
	kunal := person{
		firstName: "Kunal",
		lastName:  "Avghade",
		contact: contactInfo{
			email:   "kunal@example.com",
			zipCode: 12345,
		},
		age: 25,
	}
	// or
	kunal2 := person{"Kunal", "Avghade", contactInfo{email: "kunal2@example.com", zipCode: 67890}, 25}

	fmt.Printf("%+v\n", kunal)
	fmt.Printf("%+v\n", kunal2)

	var kunal3 person
	kunal3.firstName = "Kunal"
	kunal3.lastName = "Avghade"
	kunal3.contact = contactInfo{
		email:   "kunal3@example.com",
		zipCode: 11111,
	}
	kunal3.age = 25

	fmt.Printf("%+v\n", kunal3)
}
