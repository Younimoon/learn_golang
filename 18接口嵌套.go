package main
import "fmt"
func main(){
	// 定义一个接口不可使用这个方法（这个是定义结构体）
	// cat := Cat{}

	// 接口嵌套
	var cat a
	cat = Cat{}
	cat.test1()
	var cat1 c
	cat1 = Cat{}
	cat1.test3()
	cat1.test2()

	// 接口断言

	// 使用方法：instance := 接口对象.(实际类型) // 不安全
	// 方法二 instance,ok := 接口对象.(实际类型)
	// ok 就是判断标志（true,false）
	cicle := Cicle{10,20,30}
	fmt.Println(cicle.cicle())
	var cir circle
	cir = Cicle{1,2,3}
	fmt.Println(cir.cicle())
	
	getType(cir)
	getType(cicle)
	getType(Cat)
	// 传入一个接口的实现和结构体的实现都能进行调用
	// 为了去区分是哪个结构体或者接口引入了断言（目前的理解）
	// fmt.Println(cir)
	// 函数传参判断参数的类型
	

}
func getType(p circle){
		if ins,ok := p.(Cicle);ok{
			fmt.Println(ins,"接口调用")
			return 
		}
		fmt.Println("方法被调用啦")
	}
type a interface{
	test1()
}
type b interface{
	test2()
}
type c interface{
	a
	b
	test3()
}

type Cat struct{

}
func (cat Cat)test1(){
	fmt.Println("test1")
}
func (cat Cat)test2(){
	fmt.Println("test2")
}
func (cat Cat)test3(){
	fmt.Println("test3")
}

type Cicle struct{
	a,b,c int
}

//实现的方法中存在返回值，接口中必须写返回值
type circle interface{
	cicle() int
}

func (p Cicle)cicle()int{
	return 20
}
