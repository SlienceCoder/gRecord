package main

import (
	"fmt"
)

// 1.正常的for
func testFor1() {
	var i int

	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("final--%d\n", i)
}

// 2.带break的for

func testFor2() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("index--%d\n", i)
		if i > 5 {
			break
		}
	}
	fmt.Printf("final==%d\n", i)
}

// 3.continue的for
func testFor3() {
	var i int
	for i = 0; i < 10; i++ {

		if i%2 == 0 {
			continue
		}
		fmt.Printf("inde--%d\n", i)

	}
	fmt.Printf("final--%d\n", i)
}

// 4.省略第一个语句
func testFor4() {
	var i int = 1
	for ; i < 10; i++ {
		fmt.Printf("index--%d\n", i)
	}
	fmt.Printf("final--%d\n", i)
}

// 5.省略第一个和第三个
func testFor5() {
	var i int
	for i < 11 {
		fmt.Printf("index --%d\n", i)
		i++
	}
	fmt.Printf("final--%d", i)
}

// 6.多重赋值
func mutil_value() (a, b, c int) {
	a, b, c = 1, 2, 3
	return
}

// 7.两个变量的for
func testFor6() {
	for a, b := 1, 2; a < 10 && b < 20; a, b = a+1, b+1 {
		fmt.Printf("a=%d,b=%d\n", a, b)
	}
	fmt.Println("结束----------------")
}

// 8.死循环
func die_for() {
	for {
		fmt.Println("haha ")
	}
}

func main() {
	// testFor1()
	// testFor2()
	// testFor3()
	// testFor4()
	// testFor5()
	// testFor6()
	// fmt.Println(mutil_value())
	die_for()
}
