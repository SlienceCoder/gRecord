package main

import(
	"fmt"
)
// 1.a,*b 给*b赋值
func testPoint1(){
	var a int 
	a = 100
	fmt.Printf("--%v  %p\n",a,&a)

	var b *int
	b = &a
	fmt.Printf("--%v  %p\n",*b,b)
}

// 2.a,*b改变*b值，打印a,*b
func testPoint2(){
	var a int 
	a = 99
	fmt.Println(a)

	var b *int
	b = &a
	*b = 7777
	fmt.Println(a,b)
}

// 3.参数为int指针
func modifyInt(a *int){
	*a = 999
}

func testPoint3(){
	var a = 88
	fmt.Printf("--%T  %v",a,a)
	modifyInt(&a)
	fmt.Printf("--%T  %v",a,a)
}

// 4.参数为指针数组

func modifyArr(a *[3]int){
	// a[0]=889
	(*a)[0]=333 // 语法糖所以a[0]也能直接调用

	/*
		结构体里面的指针类型也可以直接调用
	*/
}

func testPoint4(){
	var arr [3]int = [3]int{1,2,3}
	fmt.Println(arr)
	modifyArr(&arr)
	fmt.Println(arr)
}

// 5.new创建int和数组
func testPoint5(){
	var a *int = new(int)
	*a = 22
	fmt.Printf("---%T  %p %v\n",a,a,*a)

	var arr *[3]int = new([3]int)
	(*arr)[0] = 99
	arr[1]=88
	arr[2]=77
	fmt.Printf("--%T %p %v\n",arr,arr,*arr)


}
// a,*b ,*c
func testPoint6(){
	var a int = 9
	var b *int = &a
	var c *int = b

	*c = 8888

	fmt.Println(a,*b,*c)
}


func main(){
// testPoint1()
// testPoint2()
// testPoint3()
// testPoint4()
// testPoint5()
testPoint6()


}