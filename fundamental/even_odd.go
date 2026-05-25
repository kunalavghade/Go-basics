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
