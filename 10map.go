// map 引用类型
// map相当于对象
// 	创建一个类型为map的变量
// 	var map1 = map[type1]type2
// 	type1是键的类型  type2是值的类型
// 	make方法创建 var map2 = make(map[type1]type2)
	
// 	未进行初始化的话默认为nil相当于null
// 	创建并初始化一个map
// 	var map3 = map[string]int{"age":10}

// 	对初始化的map进行操作
// 	map["age"] = 2
// 	获取对应键的值
// 	value ：= map["age"]
// 	 如果 map 中不存在该键，则返回值类型对应的默认值
// 	判断返回值是默认值的话，我们定义的键中是否真的存在该键
// value:ok :=map["age"] value代表值,ok代表是否存在该键，存在就是true
package main
import "fmt"
func main(){
	var map1 = map[int] string{1:"10"}
	map2 := map[string]string{"name":"huang"}
	var map3 = make(map[string]string)
	map3["sex"] = "男"
	fmt.Println(map1,map2,map3)
}