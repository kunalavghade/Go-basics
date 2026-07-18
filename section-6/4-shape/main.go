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
