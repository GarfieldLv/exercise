package main

import (
	"fmt"
)

type color byte

const (
	black color = iota
	red
	yellow
	green
)

func test(c color) {
	println(c)
}

func to2(i int) {
	fmt.Printf("to2: %b\n", i)
}

func to8(i int) {
	fmt.Printf("to16: %#o\n", i)
}

func to16(i int) {
	fmt.Printf("to16: %#x\n", i)
}

//func main() {
//	test(red)
//	test(100)
//	x := 2
//	test(color(x))
//	to2(100)
//	to8(100)
//	to16(100)
//	c := "1100100"
//	a, _ := strconv.ParseInt(c, 2, 32)
//	println("a: ", a)
//}
