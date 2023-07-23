package main
import "fmt"

// 一般结构体类型声明是写再函数外部

// 这种最后一个不用加逗号
// person={a:""}
//这种最后一个需要加逗号
// person = {
	// a:"",
// }
type Address struct{
		city,state string
	}
type Person struct{
	name string
	age  int
	address Address
}

// 指针类型的结构体
type Person1 struct{
	name string
	age  int
	address *Address
}
func main(){
	// 函数内部执行从上到下
	// 结构体嵌套
	
	// ------------
	person := Person{
		name:"huang",
		age:19,
		address:Address{
			city:"万州",
			state:"与东花园",
		},
	}
	// -----------
	person1 := Person{}
	person1.name="威鹏"
	person1.age =20
	person1.address = Address{}
	person1.address.city = "aaaaaaaaaa"
	person1.address.state = "bbbb"
	// -------------
	address1 := Address{
		city:"huang",
		state:"aaasdadsadas",
	}
	person3 :=Person {
		name:"bt",
		age:20,
		address:address1,
	} 

	fmt.Println(person3)
	fmt.Println(person.address,person1.address)

	fmt.Println("+++++++++++++++++++++++++++++++")

	address2 := Address{
		city:"zhongguo",
		state:"duduuu",
	}
	person2 := Person1{
		name:"huanggou",
		age:20,
		address:&address2,
	}
	fmt.Println(person2)
}