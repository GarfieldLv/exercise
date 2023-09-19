package main

import "fmt"

func main() {
	/*age := 19

	if age > 18 {
		fmt.Println("成年了！")
	} else {
		fmt.Printf("您還未成年！")
	}
	*/

	if age := 19; age > 18 {
		fmt.Println("成年了！")
	} else {
		fmt.Printf("您還未成年！")
	}

}
