package main
import "fmt"
//结构体
// 首字母大写是公有结构体，在其他模块调包时可以调用
// 首字母小写代表私有
// 是一个具有相同或不同类型的数据结构的数据集合

//定义一个结构体

type Book struct{
	name string
	btype string 
}
func main(){
	//声明一个结构体
	// 加了{}代表初始化，不加代表一种类型
	//第一种
	// const book = Book{"huabg", "nan"}

	// 第二种
	// var book1  Book
	// book1.name = "黄"
	// book1.btype = "男"

	// 第三种
	// book2 := Book{}
	// 或者 var book3 Book

	// 第四种
	book3 := Book{name:"huang",btype:"123"}

	book4 := book3
	book4.name ="wei"
	// 修改了book4，book3的值未修改说明结构体是值传递，基本数据类型
	fmt.Println(book3,book4)
	printBook(book3)
	
	// 定义结构体指针
	p := &book4
	//通过其他变量来修改基本数据类型的方法
	(*p).name = "peng"
	// 如果我们知道这个变量指向的是某个基本数据类型的地址可以省略*
	p.name = "bo"
	fmt.Printf("%p\n",p)
	fmt.Println(book4)

	// 匿名结构体
	teacher := struct{
		name string
		age int
	}{
		name:"huang",
		age:20,
	}
	fmt.Println(teacher.name)

	// 匿名字段(缺点：字段类型不能重复)

	type woker struct{
		int
		string
	}
	woker1 := woker{11,"huang"}
	fmt.Println(woker1)

	// 将结构体作为匿名字段(又叫提升字段，好处：可直接访问匿名字段的结构体)
	// 有点继承的味道
	type Person struct{
		name string
		age int
	}
	type Person1 struct{
		sex string
		Person
	}
	var person Person1
	person.name = "'younimoon'"
	// person := Person1{}
	// person.name = "younimoon"
	fmt.Println(person.name)

}
func printBook(book Book){
	fmt.Println(book)
}