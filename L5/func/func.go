package main

import(
	"fmt"
)
// 1.普通
func sayHello(){
	fmt.Println("hello")
}

// 2.多参数+返回值
func testFunc1(a int,b int,args...int)int{
	fmt.Println(a,b,"|",args)
	fmt.Printf("---%T\n",args)
	// var s []int= []int(args)
	var sum int
	for _,v := range args{
		sum += v
	}
	sum += a
	sum += b
	return sum
}

// 3.多参数+多返回值(可变参数必须是参数里面的最后一位)
func testFunc2(a int ,b string ,args...string) (s string){
	fmt.Println("haha ==========",args[0])
	s = s + b
	for _,v :=range args {
		s += v
	}
	return
}

// 4.可变参数+返回值


func main(){
// sayHello()
// fmt.Println(testFunc1(1,2,3,4,5,6))
fmt.Println(testFunc2(8,"--","23","tt","=="))
var s []string = []string{"1","2","3"}
fmt.Println(testFunc2(2,"---",s...))





}