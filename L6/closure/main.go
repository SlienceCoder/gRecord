package main

import(
	"fmt"
	"strings"
	"time"
)
// 1.函数作为返回值
func Adder() func(i int)int{
	var x int
	return func(i int)int{
		x += i
		return x
	}
}
func testClosure1(){
		// f1的生命周期内，x并不会销毁，会累加,这点要注意
		f1 := Adder()
		fmt.Println(f1(3))
		fmt.Println(f1(7))
		fmt.Println(f1(6))
}

// 2.与1类似，但是在参数里面传递参数
func add(base int)func(i int)int{
	return func(i int)int{
		base += i
		return base
	}
}
func testClosure2(){
	// 这种情况和
	f2 := add(22)
	fmt.Println(f2(2))
	fmt.Println(f2(8))

}
// 3.判断是否是.jpg结尾

func endWith(s string)func(endStr string)string{
	return func(endStr string)string{
		if !strings.HasSuffix(s,endStr){
			return s + endStr
		}
		return s
	}
}

func testClosure3(){
	f3:=endWith("meitu")
	fmt.Println(f3(".jpg"))
	f4 := endWith("haha")
	fmt.Println(f4(".mp4"))
	f5 := endWith("ghygu.jpg")
	fmt.Println(f5(".jpg"))
}

// 返回多个函数
func returnMore(base int)(add func(i int)int,sub func(j int)int){
	add = func(i int)int{
		base += i
		return base
	}
	sub = func(j int)int{
		base -= j
		return base
	}
	return 
}

func testClosure4(){
	add ,sub := returnMore(10)
	fmt.Println(add(8))
	fmt.Println(sub(2))
}

// 闭包和goroute的坑(取到的都是10，打印index--10)
func testClosure5(){
	for i :=0;i < 10;i++{
		go func(){
			fmt.Printf("index--%d\n",i)
		}()
	}
time.Sleep(time.Second*2)
}

func main(){

	// testClosure2()
	// testClosure3()
	// testClosure4()
	testClosure5()
}