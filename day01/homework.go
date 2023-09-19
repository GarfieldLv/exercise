package main

import "fmt"

// Go 语言函数外的语句必须要以关键字开头
var name = "阿吧阿巴"
var age int

const (
	NUM = 100
)

// 打印九九乘法表
func main() {
	// 函数内部定义的变量必须要使用
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
