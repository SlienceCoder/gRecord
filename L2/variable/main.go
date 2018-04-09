package main

import(
	"fmt"
)

func main(){
// 变量可以只声明不初始化.所有的未初始化的编译器默认值 int->0 string->"" bool ->false 
	// var a string
	// fmt.Println(a)

	var (
	s int
	b string
	c bool

	)
	fmt.Println(s,b,c)

	// 局部变量可以简短的声明方式
	d := "hah"
	fmt.Println("编译器自动声明",d)

}