// defer:延迟语句,延迟一个函数或方法的调用
// 多个defer时 先defer的后执行
package main
import "fmt"
func main(){
	i := 1
	defer fmt.Println(i)
	i++
	return
}
