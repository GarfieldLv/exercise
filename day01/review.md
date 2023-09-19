## Go 安装
$GOPATH: 自己写 go 代码的工作区，保存 go 代码。

$GOROOT：安装 go 的地址

`go env`

`GOPATH/bin` 需要添加到环境变量
，`go install` 命名会讲生成的二进制可执行文件拷贝到 `GOPATH/bin`

## Go 命令
`go build`: compile go

`go build -o "xx.exe"`: compile into xx.exe

`go run main.go`: execute main. Go

`go install`: compile && copy

## Go 语言文件基础语法
存放 Go 源代码的文件后缀名是 `.go`

文件第一行 `package` 关键字声明包名

如果要编译可执行的文件，必须要有 main 包和 main 函数

Go 语言函数外的语句必须要以关键字开头

函数内部定义的变量必须要使用

## 变量
三种变量声明的方式:
1. `var name string`
2. `var name = "hahhh"`
3. 函数内部专属：`name := "babab"`

匿名变量（哑元变量）：当有些数据必须要用变量接收但是又不使用它的时候就可以用`_`来接收

## 常量
`const PI = 3.1415`

`const ERROR_NOT_FOUND = 404`

iota：实现枚举

两个要点：
1. `iota` 在 const 关键字出现时将被重置为0
2. const 每新增一行常量声明，iota 累加 1

## 流程控制
if
```text
if 条件 {
    执行语句
}else if 条件 {
    执行语句
}else{
    执行语句
}
```

for 

for range
## 基本数据类型
### 整型
    
    无符号整型：uint8（11111111）、uint16、uint32、uint64
    
    带符号整型：int8、int16、int32、int64
    
    uint 和 int：具体是32位还是64位看操作系统

    uintptr: 表示指针

### 其他进制数
Go 语言中不能直接定义一个二进制数，可以直接定义八进制数和十六进制

    // 八进制数
    var n1 = 07777
    
    // 十六进制
    var n2= 0xff
### 浮点型 

`float64`和`float32`
默认是 `float64`

### 复数

`complex128`和`complex64`

### 布尔值

`true`和`flase`

不能和其他的类型做类型转换

### 字符串
常用方法

字符串不能修改

### byte 和 rune 类型
int8、int32
都属于类型别名

### 字符串、字符、字节都是什么？
字符串：双引号包裹的是字符串

字符：单引号包裹的是字符，单个字母、单个符号、单个文字

字节：1 byte = 8 bit

Go 语言中字符串都是 UTF-8 编码，UTF-8 编码中一个常用汉字一般占用3个字节。