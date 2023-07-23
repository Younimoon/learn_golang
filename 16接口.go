package main
import "fmt"
// 接口
// 接口内部只有方法的定义，不存在方式的实现
func main(){
	// 实现的第一种方法
	m1 := Mouse{"鼠标"}
	// m2 := FlassDisk{"键盘"}
	test(m1)
	// 实现就是看传进的参数中是否存在接口中的方法，且结构体的方法必须和接口中的方法一样多

	// 第二种方法(不声明一个接口函数)
	var usb USB
	usb = m1
	// 只能访问到m1中对应的接口方法，其他内容无权访问
	usb.start()

	// 总结：就是将需要使用的接口方法与之对应的结构体实例化后复制给初始化的接口

}
type USB interface {
	start()
	end()
}
type Mouse struct{
	name string
}
type FlassDisk struct{
	name string
}

func (p Mouse)start(){
	fmt.Println(p.name,"开始了")
}
func (p Mouse)end(){
	fmt.Println(p.name,"结束了")
}
func (p FlassDisk)start(){
	fmt.Println(p.name,"开始了")
}
func (p FlassDisk)end(){
	fmt.Println(p.name,"结束了")
}

// 定义一个接口
// 实现的接口必须是USB类型的
func test(usb USB){
	usb.start()
	// usb.end()
}
// 接口其实就是一个数据类型，
// 方法其实是存在于结构体当中，
// 但是我们想再定义的接口类型中调用就必须将结构体的初始化复制给接口的变量
// 可调用的方法就是接口中的方法
// 接口中的方法只有全部存在于结构体才可进行调用
