package main
// import "fmt"
import (
	"fmt"
)
//float32占用4字节
//float64占用8字节
//float由符号位+指数位+尾数位
//E代表10，E-2代表10的-2次方
func main(){
	// 定义浮点数
	var num1 float32 = 3.14
	fmt.Println(num1)
	
	var num2 float32 = -314E-2
	fmt.Println(num2)

	var num3 float32 = 314E+2
	fmt.Println(num3)
}