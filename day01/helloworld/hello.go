// 每个 go 程序都必须要有一个 package
// 每个 go 程序都是 .go 结尾
package main

// 导入标准包 fmt，格式化输出
import "fmt"

// 主函数，所有函数必须使用 func 开头
// 函数返回值是放到参数后面，不放在 func 之前
// 函数花括号必须与函数名同一行
func hello() {

	// 不需要分号结尾
	fmt.Println("Hello, World!")
}
