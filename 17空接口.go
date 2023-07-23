package main
import "fmt"
func main(){
	// 空接口 interface{}
	//空接口可以存储任意数据类型
	// var cat Cat = Cat{
	// 	color:"red",
	// }
	
	// 定义一个空接口类型的变量
	var a1 any
	a1 = Person{
		name:"huang",
		age:12,
	}
	var a2 any = "男"
	test(a2)
	fmt.Println(a1)
	// 定义一个值是任意类型的map
	map1 := make(map[string]interface{})
	map1["name"] = "huang"
	map1["age"] = 10
	fmt.Println(map1)
	// 切片里面的数据是任意数据类型
	splice1 := make([]interface{},0,10)
	splice1 = append(splice1,a1,map1)
	fmt.Println(splice1)
}

type any interface{

}
type Cat struct{
	color string

}

type Person struct{
	name string
	age int
}
// 调用函数传输参数的数据类型
func test(p any){
	fmt.Println(p)
}
// 更简单的方法
func test1(p interface{}){
	fmt.Println(p)
}