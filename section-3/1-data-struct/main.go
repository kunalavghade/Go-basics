package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type job struct {
	title string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	age       int
	job
}

func (p person) printInfo() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newName string) string {
	(*p).firstName = newName
	return p.firstName
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
	kunal2 := person{"Kunal", "Avghade", contactInfo{email: "kunal2@example.com", zipCode: 67890}, 25, job{title: "Software Engineer"}}

	kunal.printInfo()

	kunal2.updateName("Coder")
	kunal2.printInfo()

	var kunal3 person
	kunal3.firstName = "Kunal"
	kunal3.lastName = "Avghade"
	kunal3.contact = contactInfo{
		email:   "kunal3@example.com",
		zipCode: 11111,
	}
	kunal3.age = 25
	kunal3.job = job{title: "Product Manager"}
	kunal3.printInfo()
	kunalPtr := &kunal3
	kunalPtr.updateName("Kunal Updated")
	kunal3.printInfo()
}
