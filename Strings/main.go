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

	fmt.Println(strings.Index("kunalavghade@gmail.com", "@"))
	fmt.Println(strings.LastIndex("kunalavghade@gmail.com", "@"))
	fmt.Println(len("kunalavghade@gmail.com"))
	fmt.Println(strings.Repeat("a", 10))

	part := strings.Split("kunalavghade@gmail.com", "@")
	fmt.Println(part)
	
	// types
	b := strings.Builder{}
	b.Write([]byte("this is my data for builder"))
	fmt.Println(b.String())


}