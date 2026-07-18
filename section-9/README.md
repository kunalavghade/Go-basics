# Section 9

Welcome to **Section 9**. This section covers JSON Marshalling.

## Topics
- [1-marshalling](./1-marshalling/README.md)

### JSON Marshalling Example

Below is the code for the JSON marshalling topic, which demonstrates how to convert Go data structures into JSON format.

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)
// user struct represents the data we want to marshal into JSON.
// Struct tags like `json:"name"` dictate the key names in the resulting JSON.
type user struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
	IsActive bool   `json:"is_active"`
}

func main() {
	jane := user {
		Name: "jane",
		Age: 23,
		IsActive: true,
		Phone: "34567-890-567",
	}
	// json.Marshal converts a Go data structure (like our user struct) into a JSON-encoded byte slice.
	// The output is compact, without spaces or newlines.
	byteSlice, err := json.Marshal(jane)
	if err != nil{
		log.Println(err)
	}
	fmt.Println(string(byteSlice))

	// json.MarshalIndent works similarly but formats the JSON for readability.
	// The second argument is a prefix for each line (here "-"), and the third is the indent (here " ").
	byteFormated, err := json.MarshalIndent(jane, "-", " ")
	if err != nil{
		log.Println(err)
	}
	fmt.Println(string(byteFormated))
}
```

#### Explanation
1. **Struct Tags:** By adding `` `json:"key_name"` `` tags to our struct fields, we tell Go exactly what key to use in the output JSON. This ensures keys like `is_active` are used instead of the Go struct field `IsActive`.
2. **`json.Marshal`:** This function takes a Go struct and converts it into a JSON byte slice in a minified, compact format (no spaces or newlines).
3. **`json.MarshalIndent`:** This function also encodes the struct to JSON but neatly formats it with the specified prefix (`"-"`) and indentation (`" "`) to make it easily readable for humans or debugging.

---
**Previous Section:** [Section 8](../section-8/README.md)
