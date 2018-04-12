package main

import(
	"fmt"
)
// 1.初始化后下标赋值
func testFun1(){
	var a [5]int
	a[0]=9
	a[1]=8
	fmt.Println(a)

}

// 2.声明并初始化
func testArr2(){
	var a [5]int = [5]int{1,2,3,4}
	fmt.Println(a)
}

// 3.:=方式初始化
func testArr3(){
	a := [5]int{1,2,3}
	fmt.Println(a)
}

// 4.[...]方式初始化
func testArr4(){
	a := [...]int{1,2,3,5}
	fmt.Println(len(a))
}

// 5.初始化<len
func testArr5(){
	a := [6]int{1,2,3}
	fmt.Println(a)
}

// 6.下标初始化

func testArr6(){
	a := [6]int{3:9}
	fmt.Println(a)
}

// 7.初始化并赋值给同类型的另外一个变量
func testArr7(){
	a := [5]int{1,2,3}
	var b [5]int
	b = a
	fmt.Println(a)
	b[0] = 9999
	fmt.Println(a,b)
}

// 8.下标声明求长度
func testArr8(){
	a := [6]int{2:9}
	fmt.Println(len(a))
}

// 9.下标声明+遍历
func testArr9(){
	a := []int{3:88}
	for i:=0;i < len(a);i++{
		fmt.Printf("index--%d",i)
	}
	fmt.Println("final-------------")
}

// 10.for range遍历
func testArr10(){
	a := [8]int{5:90}
	for i,v := range a {
		fmt.Printf("index--%d,value--%d/n",i,v)
	}
	fmt.Println("final----------",len(a))
}

// 11.二维数组以及两种比那里方式


// 12.一个变量赋值给一个变量，修改值验证数组是值类型

// 13.数组作为参数验证数组是值传递
func modify(a [3]int){
	a[0]=999
}
func testArr13(){
a := [3]int{1,2,3}
fmt.Println(a)
modify(a)
fmt.Println(a)
}


func main(){
	// testFun1()
	// testArr2()
	// testArr3()
	// testArr4()
	// testArr6()
	// testArr7()
	// testArr8()
	testArr13()
}