package main

import "fmt"

func main() {
	//声明 slice1 是一个切片，并且初始化，默认值是1，2，3，长度是3
	//slice1 := []int{1, 2, 3}

	//声明 slice1 是一个切片，但没有分配空间
	var slice1 []int
	//slice1 = make([]int, 3) //开辟三个空间，默认值 = 0

	//声明 slice1 是一个切片，同时分配三个空间，默认值 = 0
	//var slice1 []int = make([]int,3)

	// 声明 slice1 是一个切片，同时给 slice 分配空间
	//slice1 := make([]int, 3)
	fmt.Printf("slice len is %d, slice is %v\n", len(slice1), slice1)

	// 如何判断 slice 为空
	if slice1 == nil {
		fmt.Println("slice1 is nil")
	} else {
		fmt.Println("slice1 is not nil")
	}

	slice2 := make([]int, 3, 5)
	fmt.Printf("slice2's length is %d, slice2's capacity is %d, slice2 is %v \n", len(slice2), cap(slice2), slice2)

	slice2 = append(slice2, 1)

	fmt.Printf("slice2's length is %d, slice2's capacity is %d, slice2 is %v \n", len(slice2), cap(slice2), slice2)

	slice2 = append(slice2, 2)

	fmt.Printf("slice2's length is %d, slice2's capacity is %d, slice2 is %v \n", len(slice2), cap(slice2), slice2)

	slice2 = append(slice2, 3)

	fmt.Printf("slice2's length is %d, slice2's capacity is %d, slice2 is %v \n", len(slice2), cap(slice2), slice2)

	fmt.Println("===================================================================")
	slice1 = []int{1, 2, 3}

	fmt.Printf("slice1's length is %d, slice1's capacity is %d, slice1 is %v \n", len(slice1), cap(slice1), slice1)

	slice1 = append(slice1, 1)
	fmt.Printf("slice1's length is %d, slice1's capacity is %d, slice1 is %v \n", len(slice1), cap(slice1), slice1)
	slice1 = append(slice1, 2)
	fmt.Printf("slice1's length is %d, slice1's capacity is %d, slice1 is %v \n", len(slice1), cap(slice1), slice1)

	// [0, 2) 浅拷贝，指向同一个地址的数值，数值发生变化时，两个切片的值均发生变化
	slice4 := slice1[0:2]

	fmt.Printf("slice4: %v\n", slice4)
	slice1[1] = 100
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice4: %v\n", slice4)

	// 深拷贝，只拷贝数值，不指向同一个地址
	slice5 := make([]int, 3)
	copy(slice5, slice1)
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice5: %v\n", slice5)

}
