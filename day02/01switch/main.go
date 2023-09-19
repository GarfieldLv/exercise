package main

import "fmt"

// swicth demo
func main() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇哥")

	case 2:
		fmt.Println("食指")

	case 3:
		fmt.Println("中指")

	case 4:
		fmt.Println("无名指")

	case 5:
		fmt.Println("小拇指")

	default:
		fmt.Println("无效输入")
	}

	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")

	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")

	default:
		fmt.Println(n)
	}

	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习")
	case age >= 25 && age < 35:
		fmt.Printf("好好工作\n")
	case age >= 60:
		fmt.Println("好好养老")
	default:
		fmt.Println("活着就好")
	}

	s := "a"
	switch {
	case "a" == s:
		fmt.Printf("a\n")
		fallthrough
	case "b" == s:
		fmt.Println("b")
	case "c" == s:
		fmt.Println("c")
	default:
		fmt.Println("d")
	}
}
