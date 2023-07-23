package main
// import "fmt"
import (
	"fmt"
	"unsafe"
)
//go语言的变量必须是声明赋值再使用

// 主函数自动执行 
// 外部可声明变量，但是不可进行逻辑型操作，操作一般在main函数进行
// uint（不存在符号位）和int（首位是符号位0为正1为负）
//rune 相当于int32

//go可以自动进行类型判断，不写类型也是可以的
//%t表示值的类型
//unsafe.Sizeof(e)表示变量的字节数

//全局变量
var name = "黄"
// 全局声明多个变量（推荐）
var (
	age int
	sex string
)
func main() {
	age = 1
	sex ="男"
	fmt.Println(age,sex)

	var num int = 12
	// var name = "huangweipeng "
	const a = 2
	var  b int
	//var声明b不赋值会打印默认值0，const声明就会报错

	//不用关键字声明
	c:= 3

	// 多变量的声明方式(相同类型才能这样)
	var name1,name2,name3 string
	name1,name2,name3 = "huang","wei","peng"
	fmt.Printf(name1,name2,name3)
	fmt.Println("1111111",sation1,sation2,sation3) 
    fmt.Println(num,a,b,c) 
	
	// 一次行多个变量
	//打印默认值
	var e int
	var g , v = "jack","2"
	fmt.Println(g,e,v) 
	fmt.Println(name)
	fmt.Printf("%t",e)
	fmt.Println(unsafe.Sizeof(e))
}