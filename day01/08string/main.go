package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "C:\\Users\\Garfield\\Documents\\Axure\\Libraries"
	fmt.Printf("%s\n", path) // C:\Users\Garfield\Documents\Axure\Libraries

	// 多行字符串
	s2 := `哈哈哈
	第二行
		第三行`
	fmt.Printf("%s\n", s2)
	name := "理想"
	world := "世界"
	ss1 := name + world
	fmt.Printf("len: %d\n", len(s2))
	fmt.Printf("add1: %s\n", s2+s2)
	fmt.Printf(ss1)
	fmt.Sprintf("%s%s", name, world) // 不打印，返回字符串变量
	ss2 := fmt.Sprintf("add2: " + s2)
	fmt.Printf("add2: %s\n", ss2)
	ret := strings.Split(s2, "二")
	fmt.Printf("split: %s\n", ret)
	fmt.Printf("contains 哈哈哈: %v\n", strings.Contains(s2, "哈哈哈"))
	fmt.Printf("contains 哈哈哈哈: %v\n", strings.Contains(s2, "哈哈哈哈"))
	fmt.Printf("前缀 哈哈哈: %v\n", strings.HasPrefix(s2, "哈哈哈"))
	fmt.Printf("前缀 哈哈哈哈: %v\n", strings.HasPrefix(s2, "哈哈哈哈"))
	fmt.Printf("后缀 第三行: %v\n", strings.HasSuffix(s2, "第三行"))
	fmt.Printf("后缀 第四行: %v\n", strings.HasSuffix(s2, "第四行"))
	fmt.Printf("子串出现的位置: %d\n", strings.Index(s2, "第"))
	fmt.Printf("子串最后出现的位置: %d\n", strings.LastIndex(s2, "第"))
	fmt.Printf("join 操作: %s\n", strings.Join(ret, "+"))

}
