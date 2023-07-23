package main
import "fmt"
// 指针的作用：一个指针变量指向了一个值的内存地址
// 指针的类型是 *type type的就是需要引用指针的类型

// 内存地址的转义字符：%x
func main(){
	arr := [3]int{}
	var Arr = &arr
	var b = 1
	fmt.Printf("数组的地址,%x,%p\n",&b,Arr)

	// 申明一个指针
	var p *int
	// var q int
	var a int = 10
	p = &a
	q := &a
	fmt.Printf("p:变量的地址，%x\n",&a)
	fmt.Printf("p:变量的地址,%x\n",&p)
	fmt.Printf("p:变量的地址,%d\n",p)
	fmt.Printf("q:变量的地址，%x\n",&a)
	fmt.Printf("q:变量的地址,%x\n",&p)
	fmt.Printf("q:变量的地址,%d\n",q)
	// 此时需不需要指针打印的内容是一样的


	num := 2
	p = &num
	rawNum := *p
	fmt.Println(rawNum)

	// 通过new方法创建指针
	//new(type)的type表示调用该方法访问变量的值是type类型
	num1 := new(int)
	num2 := num1
	*num2 = 10
	fmt.Printf("%p\n",num1)
	//修改num2的值但是 num1，num2存储的值是一样的，说明指针是复杂数据类型 
	fmt.Println(*num1,*num2)


	fmt.Println("--------------------------------------")

	// 数组指针 *[num]type
	//意思就是数组类型的指针和指针类型的数组
	var arr1 = [4]int{1,2,3,4}
	var p1 *[4]int
	p1 = &arr1

	fmt.Printf("%p/n",p1)
	fmt.Printf("%p/n",&p1)

	//通过p1操作数组
	(*p1)[0]=100
	//简写
	p1[0] = 200
	// 指针数组 [num]*type
	a1:= 1
	b1:= 2
	c1:= 3
	arr2 := [3]int{a1,b1,c1}
	// 就相当于arr2 ：=[3]int{1,2,3}
	// 注意 指针数组的[num]*type
	arr3 := [3]*int{&a1,&b1,&c1}
	

	fmt.Printf("%p\n",&arr2)
	// 打印的是arr2的地址，不是arr2内容的地址，arr2的地址就是arr2[0]的地址也就是数组的首地址
	fmt.Printf("%p\n",arr3[1])
	//打印的是arr3[1]存储的内容的地址
	fmt.Printf("%p\n",&b1)
	// fmt.Println("arr",&arr2[1])

	arr2[0] =100
	//数组里面的数值只是单纯的引入的值,就像是一种值传递
	//之后修改数组的值或者修改源数据都不会对另外一个产生影响
	// fmt.Println(a1) 1
	(*arr3[2]) = 100
	// fmt.Println(c1) 100(改变了)

	b1 = 200
	fmt.Println(arr2)
	//arr2的b1不改变（与arr2无关），arr3的b1改变了但地址不变
	// 总结：就是在一个值类型a中引用一个值类型b，再去a中操作b改变的只是a的值
	// 意思就是修改值类型只有直接修改或者通过类型的地址进行修改

}

