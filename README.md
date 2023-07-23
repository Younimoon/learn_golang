go常见占位符： https://www.jianshu.com/p/66aaf908045e
go的语法规范性较强：eg: if(){   不这么写会报错

                    } else{

                    }
当我们想获取某个变量的地址时只能通过下面这种方式
    p := &book4
	fmt.Printf("%p\n",p)
    如果我们想 fmt.println(p) 打印的是&p的值

    %p 取地址
    %T 数据类型
    %v 值