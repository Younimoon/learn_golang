package main
import "fmt"

// 函数也是一种数据类型是 类型是func()
// 注意参数也是()
func main(){
	var arr = [3]int{1,2,3}
	var splice = []int{4,5,6}
	a,b:= rect(1,2)
	getSum(1,"1")
	reduce(1,2,"11","1321","23afsa ")
	fun1(splice)
	fun2(arr)
	fun3()
	// 定义一个函数类型的变量
	// fun4()
	// var fn func()
	// fn = fun4
	// fn()
	arr2 := fun4()
	fmt.Printf("%p,arr2的地址\n",&arr2)
	// fmt.Println()
	// abc :=	f()
	// fmt.Println("函数指针",abc)
	fmt.Println(a,b,splice,123)
}

// 传入可变参数
// 传入的arg是一个切片,说明一个可变参数其实是一个切片
//注意一个函数只能有一个切片，且切片只能放在形参列表最后
func getSum(a int , b string,arg ... int){
	fmt.Println(a,b,arg)
}
// 下面这种传参只适用于同一类型传参
func reduce(a,b int,arg...string){
	fmt.Println(a,b,arg)
}

// 重：调用函数时将实参传给形参时相当于一次赋值（和基本数据类型和复杂数据类型扯上干系）
// 传的是值传递类型改变形参，调用的实参不改变，引用类型的话改变形参实参会改变
func fun1(s1 [] int){
	s1[0]= 100
	fmt.Println(s1)
}
func fun2(srr [3]int){
	srr[0] = 200
	
	// rect(1,2)
	fmt.Println(srr)
}

// 支持多返回值
func rect(width,height float64)(float64,float64){
	var a =width*height
	var b = height+width
	c,d := rect2(20,30)
	// 舍弃运算符_（之前也接触过） 
	// 我们指向获取一个参数进行操作
	f,_ := rect2(40,50)
	fmt.Println(a,b,c,d,f)
	return a,b
}
// 可以直接在返回的类型约束中定义返回的参数
// 函数返回的结果必须和函数定义的一致：类型 个数 顺序
func rect2(a,b int)(c,d int){
	c = a 
	d = b

	return c,d
}

func fun3(){
	for i:=0 ; i<10 ; i++{
		if i== 5  {
			// break
			return
		} else {
			fmt.Println(i)
		}
	}
}

// 指针函数
// 一个函数的返回值是指针类型数据
func fun4()(a [3]int){
	arr:=a
	fmt.Println("arrrrrrr",arr)
	return arr
	// return	[1,2,3,4]
}