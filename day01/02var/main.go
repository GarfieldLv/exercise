package main

import "fmt"

// 声明变量
// var name string
// var age int
// var isOK bool

// 批量声明变量
var (
	name string //""
	age  int    // 0
	isOk bool   // false
)

func main() {
	name = "理想"
	age = 16
	isOk = true
	// Go 语言中声明变量必须要使用，否则会报错
	fmt.Print(isOk)
	fmt.Printf(", name: %s", name)
	fmt.Println(", age:", age)

	// 声明变量同时赋值
	var name2 string = "whb"
	fmt.Println(name2)

	// 类型推导：根据值判断该变量是什么类型
	var name3 = "zxf"
	fmt.Println(name3)

	// 简短变量声明
	name4 := "xxxx"
	fmt.Println(name4)

	// 匿名变量：在使用多重赋值时，如果想要忽略某个值时可以使用匿名变量“_”。
	// 匿名变量不占用命名空间，不会分配内存，不存在重复声明
	x, _ := foo()
	fmt.Println(x)

}

func foo() (int, string) {
	return 10, "nothing"
}
