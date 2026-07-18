# 1. Strings

This topic covers the `strings` standard library package for string manipulation.

### Code Example (`main.go`)

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "abc"

	// builtin functins
	s2 := strings.Clone(s1)
	fmt.Println(s1, s2)

	fmt.Println(strings.ToLower("HEllO"))
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.ToTitle("hello"))
	fmt.Println(strings.TrimSpace("  hello im am.   kunal   "))

	fmt.Println(strings.HasSuffix("kunalavghade@gmail.com", ".com"))
	fmt.Println(strings.HasPrefix("kunalavghade@gmail.com", "kunalavghade"))
	fmt.Println(strings.Contains("kunalavghade@gmail.com", "@"))
	fmt.Println(strings.Count("aaaaaaa", "a"))

	part := strings.Split("kunalavghade@gmail.com", "@")
	fmt.Println(part)
	
	// types
	b := strings.Builder{}
	b.Write([]byte("this is my data for builder"))
	fmt.Println(b.String())
}
```

#### Explanation
1. **`strings` functions**: Go's `strings` package provides utility functions for manipulating text like `ToLower`, `ToUpper`, `TrimSpace`, `Split`, `Contains`, and more.
2. **Strings Builder**: `strings.Builder` is used to efficiently build strings using memory buffers, minimizing memory allocations compared to simple string concatenation with `+`. It implements `Write` and can return the final result with `.String()`.

## Run
```bash
go run main.go
```
