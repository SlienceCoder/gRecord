package main

import(
	"fmt"
)
// 1.切片空判断
func testSlice1(){
	var a []int
	if a == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

// 2.b从a数组切片，打印b
func testSlice2(){
	a := [...]int{1,2,3,4,5,6,7,8,9}
	b := a[2:6]
	b = append(b,888) // append会替换掉底层数组对应的值
	fmt.Println(a,b)
	b[0]=777
	fmt.Println(a,b)
	fmt.Println(b,len(b),cap(b))
	c := b[1:6] 
	// c := make([]int,3,5)
	// fmt.Println(len(c),cap(c))
	// fmt.Printf("---%p  %p  %p  %p\n",&a[0],&a,b,c)
	fmt.Println(c)
	for i:=0;i<12;i++{
		c = append(c,7)
	} 
	fmt.Println(c)
	// fmt.Printf("---%p  %p  %p\n",&a,b,c)


	fmt.Println("-----------------------------")
	// var a []int
	// a = make([]int, 5, 10)
	// a[0] = 10
	// //a[1] = 20
	// fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	// a = append(a, 11)
	// fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))

	// for i := 0; i < 8; i++ {
	// 	a = append(a, i)
	// 	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	// }
	// //观察切片的扩容操作，扩容的策略是翻倍扩容
	// a = append(a, 1000)
	// fmt.Printf("扩容之后的地址：a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
}

// 3.:=切片操作

func main(){
	// testSlice1()
	testSlice2()
}