package main
import "fmt"
func main(){
	arr2 := [3]string{"1","2"}
	var splice [] int // 空数组
	// 初始化一个具体长度切片
	var splice1 []int = make([]int ,5)
	// 简写,第二个参数是长度，第三个是容量
	splice2 := make([]int,3,10)

	//对切片元素初始化
	splice3 := [] int {1,2,3}
	// 数组中截切成切片
	splice4 := arr2[1:2]
	fmt.Println(splice,splice1,splice2,splice3,splice4)
}