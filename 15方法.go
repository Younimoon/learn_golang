package main
import "fmt"
// 方法的语法
// func (接收者)(形参列表)(返回类型){}
// 方法的接收者可以再方法内部访问到
// 方法名是可以重复的，调用方法是看调用者
func main(){
	woker := Woker{"wei",18,"男"}
	woker.towork()
	fmt.Println(woker.name)
	// 子类直接调用父类方法
	// 初始化匿名结构体的方法
	// 1
	// son:=Son{Father{"huang",19},"男"}
	// 2
	son :=Son{
		sex:"nan",
	}
	son.name = "huang"
	son.age = 20
	// 3
	p := Father{
		name:"wei",
		age:22,
	}
	son = Son{
		Father:p,
		sex:"nan",
	}

	son.tosed()
}
type Woker struct{
	name string
	age int
	sex string
}
// p就相当于实例化的worker复制的形参名
// 必须初始化一个结构体然后调用的时候将调用的内容复制给p （值传递），当我修改p的内容原初始化对象不改变
func (p Woker)towork(){
	p.name = "huang"
	fmt.Println(p.name)
}
// 子类父类的概念：结构体（子类）中嵌套一个匿名的结构体（父类）作为字段（提升字段）
//匿名结构体存在一个字段，声明的结构体（子类）可直接调用该（父类）方法
type Father struct{
	name string
	age int
}
type Son struct{
	Father
	sex string
}
// 该方法是父类的方法
func (p Father)tosed(){
	fmt.Println(p.age)
}
// 重写父类的方法（子类定义一个方法名和父类一样的，后续调用子类调用会子类的方法）

func (p Son)tosed(){
	fmt.Println(p.name)
}