package main

import (
	"fmt"
	"time"
)

// 从 goroutine
func i() {
	i := 0
	for {
		fmt.Println("new i: ", i)
		i++
		time.Sleep(1 * time.Second)
	}
}

// 主 goroutine，停止之后从 goroutine 也不会继续执行
func main() {
	go i()
	i := 0
	for {
		fmt.Println("main i: ", i)
		i++
		time.Sleep(1 * time.Second)
	}
}
