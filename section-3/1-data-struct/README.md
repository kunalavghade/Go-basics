# 1. Data Struct

This topic covers custom data structures (structs) in Go and how to use pointers with them.

### Code Example (`main.go`)

```go
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

	kunal.printInfo()

	kunalPtr := &kunal
	kunalPtr.updateName("Kunal Updated")
	kunal.printInfo()
}
```

#### Explanation
1. **Structs**: The `struct` keyword lets you create custom data types by grouping together fields of different types (e.g. `person` holds strings, ints, and other structs).
2. **Embedding**: The `person` struct embeds the `job` struct without explicitly giving it a field name. This is called embedded or anonymous fields.
3. **Pointers**: `updateName` uses a pointer receiver `*person` so that it can mutate the actual `person` object instead of a copy. `(*p).firstName` dereferences the pointer to update the value.

## Run
```bash
go run main.go
```
