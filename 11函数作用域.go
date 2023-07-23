package main
import "fmt"
// 在同一作用域下不可申明两个相同名的变量，但是在其子作用域下可以
// 子作用域修改父作用域的变量是永久性的
func main(){
	var n = 10
	if(n == 10){
		n = 20
		sum:=getSum(5)
		fmt.Println(sum)
	}
	// getSum
	// var n = 30
	fmt.Println(n)
}

// 递归函数
// 求和100
// var sum = 0
func getSum(n int)(int){
	if(n ==1){
		// return getSum(n-1)+n
		return 1
	}
	return getSum(n-1)+n
	
}