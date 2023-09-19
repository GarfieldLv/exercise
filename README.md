# go
## 一、初识

### windows 安装
官网下载 go1.16.13.windows-386.msi，直接安装不需要自定义环境变量，使用 goland 非常方便。
### 环境变量

go env
```
set GO111MODULE=off
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\Garfield\AppData\Local\go-build
set GOENV=C:\Users\Garfield\AppData\Roaming\go\env
set GOEXE=.exe
set GOFLAGS=
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOINSECURE=
set GOMODCACHE=D:\project\go-exercise\pkg\mod
set GONOPROXY=
set GONOSUMDB=
set GOOS=windows
# go 工程目录
set GOPATH=D:\project\go-exercise
set GOPRIVATE=
set GOPROXY=https://goproxy.cn
# go 安装目录
set GOROOT=D:\Go
set GOSUMDB=sum.golang.org
set GOTMPDIR=
set GOTOOLDIR=D:\Go\pkg\tool\windows_amd64
set GCCGO=gccgo
set AR=ar
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set GOMOD=
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 -mthreads -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=C:\Users\Garfield\AppData\Local\Temp\go-build280570006=/tmp/go-build -gno-record-gcc-switches
```

### go 语言特点
1. 没有头文件概念
2. 强类型的语言，编译型语言（python 是弱类型，解释型语言）
3. go 语言应用程序，在运行时时不需要依赖外部库的

    1）执行时所需要的库都打包到程序中

    2）go 程序会比较大
    
    3）未使用到 import 的包，会报错

    4）平台无关，需要两个环境变量来控制：
        
        （1） GOOS：设定运行的平台
            mac: SET GOOS=darwin
            linux: SET GOOS=linux
            windows: SET GOOS=windows

        (2) GOARCH: 目标平台的体系构架
            386: SET GOARCH=386
            amd64: SET GOARCH=amd64
            arm: SET GOARCH=arm

    最好在一开始禁用 CGO：SET CGO_ENABLED=0再设置前面两个参数。
    
### go command
1. 项目目录下执行 go build；
2. 在其他路径下执行 go build，需要在后面加上项目的路径，路径从 GOPATH/src 后开始邪气，编译之后的可执行文件即出现在当前目录下；
3. 编译 .go 文件，-o 指定生成文件
```shell
    go build -o test.exe main.go xxx.go
    go build *.go
```

4. go run：执行脚本文件一样执行 go 代码
5. go install
   1） 先编译，得到可执行文件
   2）将可执行文件拷贝到 /GOPATH/bin，在任何地方都可执行

## 二、语法
### 变量声明
Go 语言的变量声明格式为
```shell
var 变量名 变量类型
```
行尾不需要分号。


### 批量声明
```go
package main

var (
	name string
	age  int
	isOk bool
)
```
### 变量的初始化
注意：Go 语言中未使用未声明的变量会报错
### 类型推导
### 简短变量声明
### 匿名变量

注意：
1. 函数外的每个语句都必须以关键字开始（var、func、const 等）
2. := 不能在函数外使用
3. _ 多用于占位，表示忽略

### 常量 const
```go
 package main 
   // PI 常量
   const PI = 3.1415926
   
   // 批量声明
   const (
   // STATUS = 200
   // NOTFOUND = 404
   // 批量声明时第一行赋值之后，后面几行未指定值的常量均为第一行对应的值
   n1 = 100
   n2
   n3
   )
```
#### iota
*iota* 是 go 语言的常量计数器，只能在常量表达式中使用。  

iota 在 const 关键字出现时将被重置为0，const 中每次新增一行常量声明就会计数一次。

在定义枚举时非常有用。
```go
package main
const(
	a1 = iota // 0
	b2 // 1
	b3 // 2
)

const(
    c1 = iota // 0
    c2 // 1
	_  // 2 
    c3 // 3
	c4 = 100 // 100
	c5  // 100
	c6 = iota // 4
)

const (
    d1, d2 = iota + 1, iota + 2 // 1, 2, iota 在 const 关键字出现时将被重置为0
    // 空行不会新增
    d3, d4 = iota + 1, iota + 2 // 2, 3, 新增一行常量声明 iota 增加1
)

// 定义数量级
const (
   _  = iota
   KB = 1 << (10 * iota) // 2^10 = 1024
   MB = 1 << (10 * iota)
   GB = 1 << (10 * iota)
   TB = 1 << (10 * iota)
   PB = 1 << (10 * iota)
)

```

### 基本数据类型
整型、浮点型、布尔型、字符串、数组、切片、函数、map、channel 等。

| 类型  |  具体类型   | 描述                                                   |
|:---:|:-------:|:-----------------------------------------------------|
| 整型  |  uint8  | 无符号8位整型（0到255），byte                                  |
|     | uint16  | 无符号16位整型（0到65535）                                    |
|     | uint32  | 无符号32位整型（0到2^32-1）                                   |
|     | uint64  | 无符号64位整型（0到2^64-1）                                   |
|     |  int8   | 有符号8位整型（-128到127）                                    |
|     |  int16  | 有符号16位整型（-32768到32767）                               |
|     |  int32  | 有符号32位整型（-2^16到2^16-1）                               |
|     |  int64  | 有符号64位整型（-2^32到2^32-1）                               |
|     |  uint   | 32位操作系统上是 uint32，64位操作系统上是uint64                     |
|     |   int   | 32位操作系统上是 int32，64位操作系统上是int64                       |
|     | uintptr | 无符号整型，用于存放一个指针                                       |
| 浮点数 | float32 | 3.4e38，可以使用常量 math.MaxFloat32                        |
|     | float64 | 1.8e308，可以使用常量 math.MaxFloat64                       |
| 复数  | complex64| 实部和虚部为32位，c1 = 1 + 2i                                |
|     | complex128| 实部和虚部为64位                                            |
| 布尔值 | bool | 布尔型数据只有 true 和 false 两个值；默认是 false；不能和其他类型转换，也不参与数值运算|
| 字符串 | | 字符串以 utf-8 编码（支持中文），字符串的值必须要用双引号（“”）包裹|               |


注意：
* 在使用 int 或者 uint 类型时不能假定，需要考虑不同平台上的差异；
* 获取对象长度得内建 len() 函数返回的长度可以根据不同平台的字节长度进行变化。
  * 实际使用中，切片或者 map 的数量用 int 来表示；
  * 涉及二进制传输、读写文件的结构描述时，不要使用 int 或者 uint，为了保持文件的结构不会受到不同编译平台字节长度影响。
* 浮点数默认是 float64 类型，需要声明 float32 要显式声明
```
    f2 := float32(3.14159)
  ```
* 打印浮点数可以和 %f 结合
```
   fmt.Printf("math.MaxFloat64: %.2f\n ", math.Pi)
   fmt.Printf("math.MaxFloat64: %f\n", math.Pi)
   ```
* 布尔值
```
	b1 := true
	var b2 bool
	fmt.Printf("%v\n", b1) // true
	fmt.Printf("%v\n", b2) // false
```
* 区分字符串、字符、字节

单独的字母、汉字和符号表示一个字符

字节：1字节 = 8 Bit（8个二进制位）

1个字符 'A' = 1 个字节

1个 utf8 编码的汉字，一般占用3个字节


#### 整型应用--八进制 & 十六进制
```go
package main

import "fmt"

// 整型

func main() {
	// 十进制
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) // 十进制数转成二进制
	fmt.Printf("%o\n", i1) // 十进制数转成八进制
	fmt.Printf("%x\n", i1) // 十进制数转成十六进制
	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)

	// 十六进制
	i3 := 0x1234567
	// 查看变量的类型
	fmt.Printf("%T\n", i3)
	fmt.Printf("%d\n", i3)

	i4 := int8(9) // 明确声明 int8 变量类型，否则默认 int 类型
	fmt.Printf("%T\n", i4)
}
```

#### 字符串
1. 转义字符
2. 多行字符串需要用 `` 括起来，包含其中的空格
```
   	// 多行字符串
	s2 := `
		哈哈哈
	第二行
		第三行
	`
	fmt.Printf("%s\n", s2)
```
3. 字符串的常用操作

|                   方法                    | 描述  |
|:---------------------------------------:|:----|
|                len(str)                 | 求长度|
|            + 或者 fmt.Sprintf             | 拼接字符串|
|              strings.Split              | 分割|
|            strings.Contains             | 判断是否包含|
| strings.HasPrefix() strings.HasSuffix() | 判断字符串前缀或者后缀/
|        strings.Index() strings.LastIndex() | 子串出现的位置|
| strings.Join(a[]string, sep string) | join 操作|

4. byte 和 rune 类型
每个组成字符串的元素被称为“字符”，可以通过遍历或者单个获取字符串元素获得字符。
Go 字符主要有两种：
* byte 类型：uint8 类型，代表了 ASCII 码的一个字符
* rune 类型：int32 类型，代表一个 UTF-8 类型

当需要处理中文等复合字符时，则需要 rune 类型。
```go
package main

import "fmt"

func main() {
	s := "Hello, 沙河"
	// len() byte 字节的数量
	n := len(s)
	fmt.Println(n)

	for _, c := range s { 
		fmt.Printf("%v(%c)\n", c, c)
	}

   c1 := "红"
   c2 := '红'                            //rune(int32)
   fmt.Printf("c1:%T\nc2:%T\n", c1, c2) // string, int32
   fmt.Printf("%d\n", c2)
}
```

```text
13
72(H)    
101(e)   
108(l)   
108(l)   
111(o)   
44(,)    
32( )    
27801(沙)
27827(河)
c1:string
c2:int32 
32418    

Process finished with the exit code 0
```

5. 修改字符串
字符串是不能直接修改的。
需要利用 byte[] 和 rune[] 来修改
```go
package main

import "fmt"

func main() {
	//// 字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '红'
	fmt.Println(string(s3))
}
```

6. 类型转换
Go 中只有强制类型转换，没有隐式类型转换，该语法只能在两个类型之间互相支持相互转换的时候使用。
```text
T(表达式)
```
其中，T 表示要转换的类型，表达式包括变量、复杂算子和函数返回值等。
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b = 3, 4
	var c int
	// math.Sqrt() 接收的参数是 float64 类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
```
### 流程控制
if 和 for 是最常见的流程控制，switch 和 goto 是为了简化代码，属于扩展类的流程控制。

```text
for 初始语句; 条件表达式; 结束语句 {
    循环体语句
}
```

无限循环
```text
for {
    循环体语句
}
```
for 循环可以通过 break、goto、return、panic 语句强制退出循环。

#### for range（键值循环）
Go 语言中可以使用 for range 遍历数组、切片、字符串、 map 及通道（channel）。
通过 for range 遍历的返回值有以下特点：
1. 数组、切片 、字符串返回索引和值
2. map 返回键和值
3. 通道只返回通道里面的值

#### switch case
使用 `switch`语句可方便地对大量的值进行条件判断

Go 语言规定每个`switch`只能有一个`default`分支
```text
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

```

一个分支可以有多个值。多个 case 值中间使用英文逗号分隔。
```text
switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")

	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")

	default:
		fmt.Println(n)
	}
```

分支还可以使用表达式，这个时候 `switch`后面不需要跟判断变量
```text
    age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习")
	case age >= 25 && age < 35:
		fmt.Println("好好工作")
	case age >= 60:
		fmt.Println("好好养老")
	default:
		fmt.Println("活着就好")
	}
```

`fallthrough`语法可以执行满足条件的 case 的下一个 case，是为了兼容 C 语言中的 case 设计的。

```text
	s := "a"
	switch {
	case "a" == s:
		fmt.Println("a")
		fallthrough
	case "b" == s:
		fmt.Println("b")
	case "c" == s:
		fmt.Println("c")
	default:
		fmt.Println("d")
	}
```

### goto
```text
func main() {
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			if 3 == i {
				goto breakTag
			}

			fmt.Printf("%d\t", j)
		}
	}

breakTag: // label
	fmt.Println("over")
}
```

`break`语句也可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须要定义在对应的 `for`、`switch`、`select`的代码块中。
```text
BREAKDEMO:
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			if 3 == i {
				break BREAKDEMO
			}

			fmt.Printf("%d\t", j)
		}
	}

	fmt.Println("....")
```

`continue`语句可以结束当前循环，开始下一次循环的迭代过程，仅限于 `for` 循环内使用。

在`continue`语句后面添加标签时，表示开始标签对应的循环。
```text
CONTINUEDEMO:
	for i := 0; i < 10; i++ {
		for j := i; j < 10; j++ {
			if 3 == i && 5 == j {
				continue CONTINUEDEMO
			}

			fmt.Printf("%d\t", j)
		}
	}

	fmt.Println("....")
```

## TIPS
1. GoLand 中设置 shell

   Setting > Tools > Terminal 中修改 Shell path 改成 git bash 所在位置，比如 D:\Program Files\Git\bin\bash.exe.

   修改成功之后可以重启 Goland 中的 Terminal，就会变成 git bash.
   
   适用于 windows 系统。

2. vscode 中设置 shell，在 setting.json 中增加几行：
```
    "terminal.integrated.profiles.windows": {
        "gitBash": {
          "path": "D:\\Program Files\\Git\\bin\\bash.exe",
        }
      },
    "terminal.integrated.defaultProfile.windows": "gitBash",
```

