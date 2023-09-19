package main

import "fmt"

var (
	n int
	f float32
	b bool
	s string
)

func main() {
	s = "I'm a string!"
	fmt.Printf("n: %T, %d\n", n, n)
	fmt.Printf("f: %T, %f\n", f, f)
	fmt.Printf("b: %T, %v\n", b, b)
	fmt.Printf("s: %T, %s\n", s, s)
	s = "hello沙盒小丸子"
	fmt.Printf("%d\n", len(s))
}
