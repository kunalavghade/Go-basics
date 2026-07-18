# JSON Marshalling in Go

This module covers the basics of converting Go data structures (like `structs`) into JSON format, a process known as **Marshalling**.

## Key Concepts

1. **`json.Marshal`**
   - **Purpose:** Converts a Go data structure into a JSON-encoded byte slice.
   - **Usage:** `byteSlice, err := json.Marshal(data)`
   - **Output:** Compact JSON without spaces or newlines.

2. **`json.MarshalIndent`**
   - **Purpose:** Similar to `json.Marshal`, but formats the output to make it human-readable.
   - **Usage:** `byteFormated, err := json.MarshalIndent(data, prefix, indent)`
   - **Example:** `json.MarshalIndent(jane, "-", " ")` adds a `-` prefix and spaces for indentation.

3. **Struct Tags**
   - Go uses struct tags to define how struct fields map to JSON keys.
   - **Syntax:** `` `json:"key_name"` ``
   - **Example:**
     ```go
     type user struct {
         Name     string `json:"name"`
         Age      int    `json:"age"`
         Phone    string `json:"phone"`
         IsActive bool   `json:"is_active"`
     }
     ```
   - These tags ensure the resulting JSON has lowercase keys with underscores, matching the expected JSON convention rather than Go's Capitalized exported fields.

## Quick Run

```bash
go run main.go
```
