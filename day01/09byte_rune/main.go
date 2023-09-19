package main

import (
	"fmt"
	"math"
)

// byte && rune 类型
// Go 语言中为了处理非 ASCII 码类型的字符，定义了新的 rune 类型
func main() {
	s := "Hello, 沙河"
	// len() byte 字节的数量
	n := len(s)
	fmt.Println(n)

	//for i := 0; i < len(s); i++ {
	//	//fmt.Printf(s[i])
	//	fmt.Printf("%c\n", s[i])
	//}

	for _, c := range s {
		fmt.Printf("%v(%c)\n", c, c)
	}

	//// 字符串修改，字符串强制转换成了 rune 切片才能做修改
	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '红'
	fmt.Println(string(s3))

	c1 := "红"
	c2 := '红'                            //rune(int32)
	fmt.Printf("c1:%T\nc2:%T\n", c1, c2) // string, int32
	fmt.Printf("%d\n", c2)
	c3 := "H" //string
	c4 := 'H' // int32
	fmt.Printf("c3: %T\nc4: %T\n", c3, c4)
	c5 := byte('H')
	fmt.Printf("c5: %T\n", c5)

	var a, b = 3, 4
	var c int
	// math.Sqrt() 接收的参数是 float64 类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

