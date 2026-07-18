# 4. Shape

This topic covers a shape implementation using Go interfaces.

### Code Example (`main.go`)

```go
package main

type shape interface {
	area() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	println(s.area())
}

func main() {
	s := square{sideLength: 5}
	t := triangle{base: 10, height: 5}
	printArea(s)
	printArea(t)
}
```

#### Explanation
1. **`shape` Interface**: Declares a contract that any struct must have an `area() float64` method to be considered a `shape`.
2. **Struct Implementations**: Both `square` and `triangle` provide their own unique implementation of the `area()` method.
3. **Shared Functions**: Because they both satisfy the `shape` interface, they can both be passed into the `printArea(s shape)` function safely.

## Run
```bash
go run main.go
```
