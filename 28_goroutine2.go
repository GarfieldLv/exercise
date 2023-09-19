package main

import (
	"fmt"
	"time"
)

func main() {
	// 无形参的匿名函数
	//go func() {
	//	defer fmt.Println("a.defer")
	//	// return // 只取消注释此处，只会打出 a.defer
	//	func() {
	//		defer fmt.Println("b.defer")
	//		//return // 只取消注释此处，只打出 b.defer a a.defer
	//		runtime.Goexit() //只去取消注释此处, 只打出 b.defer a.defer，退出本次的 goroutine
	//		fmt.Println("b")
	//	}()
	//	fmt.Println("a")
	//}()

	// 有形参的匿名函数，参数值放在括号中，这边不能直接定义一个变量接收返回的 bool 值，需要借助 channel 才可以
	go func(a, b int) bool {
		fmt.Printf("a = %d, b = %d", a, b)
		return true
	}(10, 20)
	for {
		time.Sleep(1 * time.Second) //主 goroutine 不死亡才会执行从 goroutine

	}
}
