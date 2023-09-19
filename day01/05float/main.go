package main

import (
	"fmt"
	"math"
)

// 浮点数

func main() {
	max32 := math.MaxFloat32 // 32位 float 数
	fmt.Println("math.MaxFloat32: ", max32)
	max64 := math.MaxFloat64 // 32位 float 数
	fmt.Println("math.MaxFloat64: ", max64)
	fmt.Printf("math.MaxFloat64: %.2f\n ", math.Pi)
	fmt.Printf("math.MaxFloat64: %f\n", math.Pi)
	f1 := 1.2345
	fmt.Printf("%T\n", f1)
	f2 := float32(1.2345)
	fmt.Printf("%T\n", f2)
}
