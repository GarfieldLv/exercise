package main

import "fmt"

// PI 常量
const PI = 3.1415926

// 批量声明
const (
	// STATUS = 200
	// NOTFOUND = 404
	// 批量声明时第一行赋值之后，后面几行未指定值的常量均为第一行对应的值
	n1 = 100
	n2
	n3
)

const (
	i1 = iota // 0
	i2 = iota // 1
	i3        // 2
)

const (
	c1 = iota // 0
	c2        // 1
	_         // 2
	c3        // 3
	c4 = 100  // 100
	c5        // 100
	c6 = iota
	c7
)

const (
	d1, d2 = iota + 1, iota + 2 // 1, 2, iota 在 const 关键字出现时将被重置为0
	// 空行不会新增
	d3, d4 = iota + 1, iota + 2 // 2, 3, 新增一行常量声明 iota 增加1
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) // 2^10 = 1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {

	fmt.Printf("pi: %f\n", PI)
	fmt.Printf("n1: %d\n", n1)
	fmt.Printf("n2: %d\n", n2)
	fmt.Printf("n3: %d\n", n3)
	fmt.Printf("c5: %d\n", c5)
	fmt.Printf("c6: %d\n", c6)
	fmt.Printf("c7: %d\n", c7)
	fmt.Printf("i2: %d\n", i2)
	fmt.Printf("i3: %d\n", i3)
}
