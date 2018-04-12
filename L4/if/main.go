package main

import (
	"fmt"
)

// 1.只有两层
func testIf1() {
	num := 4
	if num == 4 {
		fmt.Println("d等于4")
	} else {
		fmt.Println("不等于4")
	}
}

// 2.多层
func testIf2() {
	num := 90
	if num >= 90 {
		fmt.Println("优秀")
	} else if num > 80 && num < 90 {
		fmt.Println("良")
	} else if num > 60 && num <= 80 {
		fmt.Println("中")
	} else {
		fmt.Println("中")
	}
}

// 3.if中有条件和判断
func testIf3() {
	if num := 6; num > 6 {
		fmt.Println(">6")
	} else {
		fmt.Println("<6")
	}
}

// 4.if包含函数求职赋值
func getNum() int {
	return 8
}
func testIf4() {
	if num := getNum(); num > 6 {
		fmt.Println(">6")
	} else {
		fmt.Println("<6")
	}
}

func main() {
	testIf1()
	testIf2()
	testIf3()
	testIf4()
}
