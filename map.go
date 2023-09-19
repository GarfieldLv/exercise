package main

import "fmt"

func main() {
	//声明 map
	var map1 map[string]string
	if map1 == nil {
		fmt.Println("这是一个空 map")
	}

	map1 = make(map[string]string, 10)
	map1["one"] = "C++"
	map1["two"] = "go"
	map1["three"] = "java"

	// 按照 hash 值排序，不是按照传入的顺序排
	fmt.Println(map1)

	map2 := make(map[int]string)
	map2[1] = "java"
	map2[2] = "C"
	map2[3] = "kotlin"
	fmt.Println(map2)

	map3 := map[string]string{
		"one":   "php",
		"two":   "objectc",
		"three": "swift",
	}
	fmt.Println(map3)

}
