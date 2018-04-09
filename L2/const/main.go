package main

import (
	"fmt"
)

func main(){
	// 之声明第一个时候，后面的两个数也相同
	const(
		a = 1
		b
		c
	)
	fmt.Println(a,b,c)

	/*
		声明不同类型的常量时，如果标识了常量的类型，需初始化值，以下这种是编译不过的
		const(
	 		a1 int = 9
	 		b1 string 
	 		c1 bool 
)
	*/
const(
	 a1 int = 9
	 b1 string = "hello"
	 c1 bool  = true
)
fmt.Printf("%d  %s  %t\n",a1,b1,c1)


// 类似像变量那样只声明不赋值会报错的 这和上面的例子一样
/*
const con string
fmt.Println(con)
*/

// 每遇到一个const,iota从0开始，依次递增
const(
	a2 = iota
	b2
	c2
	d2
)
fmt.Println(a2,b2,c2,d2)

/*
这里其实就是
a3 = 1*2^0
b3 = 1*2^1
c3 = 1*2^2
d3 = 1*2^3
*/
const(
	a3 = 1 << iota
	b3
	c3
	d3
)
fmt.Println(a3,b3,c3,d3)

// 从第一个开始记。不论是否第一个含有iota
const(
	a13 = 1
	a4 = 3 // 3
	a5 = 3 << iota // 3*2^2
	a6 // 3*2^3
	a7 = "haha" // "haha"
	a8 // "haha"
	a9 = 1 << iota // 1*2^6
)
fmt.Println(a13,a4,a5,a6,a7,a8,a9)

}