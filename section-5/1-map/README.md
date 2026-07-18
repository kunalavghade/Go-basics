# 1. Map

This topic covers the Map data structure in Go, which stores key-value pairs.

### Code Example (`main.go`)

```go
package main

func main() {
	color := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	println(color["red"])

	// Initialization using make
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2

	println(numbers["one"])

	// Deleting a key
	delete(color, "green")
	
	printMap(color)
}

func printMap(m map[string]string) {
	for key, value := range m {
		println(key, ":", value)
	}
}
```

#### Explanation
1. **Map Declaration**: Maps are declared using `map[keyType]valueType`. You can initialize them with literal data using curly braces `{}`.
2. **`make` Function**: If you don't use a literal, you must initialize the map using `make(map[keyType]valueType)` before adding elements to avoid runtime panics.
3. **`delete` Function**: You can remove key-value pairs from a map using the built-in `delete(map, key)` function. Iterating over maps with `range` returns both the key and the value.

## Run
```bash
go run main.go
```
