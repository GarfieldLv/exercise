package main

import "fmt"

func main() {
	// for range
	s := "Hello, World！"
	for i, v := range s {
		fmt.Printf("%d: %c\n", i, v)
	}
}
