package main

func dataf() []int {
	println("origin data")
	return []int{10, 20, 30}
}

//func main() {
//	data := []string{"a", "b", "v"}
//
//	for d, s := range data {
//		println(&d, &s) // 局部变量始终重复使用
//		println(d, s)
//	}
//
//	for range data { // 仅迭代，不返回，用于清除 channel 等操作
//
//	}
//
//	for d := range data {
//		println(d)
//	}
//
//	for _, s := range data {
//		println(s)
//	}
//
//	data2 := [3]int{10, 20, 30} // 这边如果是 []int{10,20,30}，后续结果是不一样的
//	for i, s := range data2 {   //拿到的只是 data 的复制品，不会变更原数值
//		if i == 0 {
//			data2[0] += 100
//			data2[1] += 200
//			data2[2] += 300
//		}
//		fmt.Printf("x: %d, data: %d\n", s, data2[i])
//	}
//
//	for i, s := range data2[:] { // 仅复制 slice，不包括底层的 array
//		if i == 0 {
//			data2[0] += 100
//			data2[1] += 200
//			data2[2] += 300
//		}
//		//if i == 1 {
//		//	data2[0] += 100
//		//	data2[1] += 200
//		//	data2[2] += 300
//		//}
//		fmt.Printf("x: %d, data: %d\n", s, data2[i]) // 第一个值先被取出来了之后再修改的，所以第一个是 110，data：210
//	}
//
//	for i, x := range dataf() { // 目标函数也只会被调用一次
//		println(i, x)
//	}
//}
