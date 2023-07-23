package main
import "fmt"
func main(){
	var arr [5]int
	// 一维数组定长不一定是固定长度是传一个定长的标识
	// 不确定长度就是[...]，确定长度就是[num]
	// 数组值传递:
	// 不使用关键字声明是需要进行赋值操作
	// 如:arr3 := [2]string{}
	// 数组长度下的某个元素没有赋值会使用默认值
	// 数组的长度固定后不可超过长度范围进行添加元素
	arr2 := [3]string{"1","2"}
	arr2[2] = "3"

	fmt.Println(arr,arr2)

	// 二维数组
	// 二维数组的长度可以不定长
	// 二维数组里面的每个一维数组是具有内存地址的（没有申明也具有内存地址）
	values :=[] []int{}
	// 下面这种在单独创建一个一维数组是会报错，一维数组的创建必须定长
	row1 := []int{1,2,3}
	row2 := []int{4,5,7}
	values = append(values,row1)
	values = append(values,row2)
	fmt.Println(values)

	var arr4 = [...]int{1,2,3,4,5}
	// 二维数组的创建
	var arr3 = [3][4]int{}
	fmt.Println(arr3,arr4)
}