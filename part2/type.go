package main

type flag byte

const (
	read flag = 1 << iota
	write
	exec
	freeze
)

//func main() {
//	f := read | exec // 按位或，至少有一个1
//	println(f)
//	println(exec)

//const v = 20
//var a byte = 10
//b := v + a
//fmt.Printf("%T,%v\n", b, b)
//
//const c float32 = 1.2
//d := c + v
//fmt.Printf("%T, %v", d, d)
//}
