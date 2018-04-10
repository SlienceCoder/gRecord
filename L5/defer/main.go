package main

import(
	"fmt"
)
// 1.多个defer同时压栈(先进后出)
func testDefer1(){
	defer fmt.Println("--1")
	defer fmt.Println("--2")
	defer fmt.Println("--3")
	fmt.Println("final....")
}

// 2.for循环里面的defer
func testDefer2(){
	for i :=0 ;i < 5;i++{
		defer fmt.Printf("index--%d\n",i)
	}
	fmt.Println("finale----")
}

// 3.defer前后改变值得情况
func testDefer3(){
	var a = 8
	defer fmt.Println(a)
	a = 999
	fmt.Println(a)
}

// 4.defer与闭包的使用,闭包访问的是地址，所以值最新
func testDefer4(){
	var a = 9
	defer func (){
		fmt.Printf("闭包和defer --%d\n",a)
	}()
		a = 888
	fmt.Println("final----",a)
}

func main(){
	// testDefer1()
	// testDefer2()
	// testDefer3()
	testDefer4()
}