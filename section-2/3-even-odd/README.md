# 3. Even Odd

This topic covers logic for finding even and odd numbers using loops and conditionals.

### Code Example (`main.go`)

```go
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		if num&1 != 0 {
			fmt.Printf("%d is odd\n", num)
		} else {
			fmt.Printf("%d is even\n", num)
		}

	}
}
```

#### Explanation
1. **Slices**: `nums := []int{...}` creates a slice (dynamic array) of integers.
2. **Range Loop**: `for _, num := range nums` iterates through the slice. The underscore `_` is a blank identifier used to ignore the index since we only need the value (`num`).
3. **Bitwise AND (`&`)**: `num & 1 != 0` is an efficient way to check if a number is odd. If the least significant bit is 1, the number is odd; otherwise, it's even. You could also use `num % 2 != 0`.

## Run
```bash
go run main.go
```
