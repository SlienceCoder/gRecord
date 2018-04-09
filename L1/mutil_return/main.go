package main

import (
	"fmt"
)
// 多返回值的返回值参数可以带参数名也可以不带
// eg：func test(a,b int) (int,int)
func test(a,b int) (c int,d int) {
	return a + b,a -b
}

func main(){
	a,b := test(30,8)
	fmt.Println(a,b)
}