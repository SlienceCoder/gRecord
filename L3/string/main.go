package main

import(
	"fmt"
)
/*
	作业：1.字符串的各种操作 2.字符串反转，回文判断
*/

func main(){
	// testString()
	fmt.Println(testhuiwen("我爱我s"))
	
	a := 10/3.0
	fmt.Println(a)
}

func testhuiwen(s string)bool{

	// var str = s
	// var a []rune = []rune(str)

	// for i := 0; i <len(a)/2;i++ {
	// 	tmp := a[len(a)-i-1]
	// 	a[len(a)-i-1] = a[i]
	// 	a[i] = tmp
	// }

	// // str = string(a)
	// fmt.Println(str)
	// fmt.Println(string(a))
	// if str ==  string(a) {
	// 	return true
	// } else {
	// 	return false
	// }

	var temp = []rune(s) // 反转后的字符串
	var a = []rune(s) // 转化为字符数组
	fmt.Println(temp)
	fmt.Println(a)
	for i := 0;i < len(a)/2;i++{
		
		temp[i]=a[len(a)-i-1]
		
	}
	fmt.Println(string(temp))
	fmt.Println(string(a))
	if string(temp)==s{
		return true
	} else {
		return false
	}
    
}


func stringRev(){

}

func testString(){
	// 字符串下标访问
	var str = "you are the one ok"
	// 遍历字符串
	for i ,v := range str {
		fmt.Printf("index=%d,对应的value=%c\n",i,v)
	}
	fmt.Println("----------------")

	// 字符数组和字符串的转换
	for i ,v := range []byte(str){
		fmt.Printf("这是---%d,%c\n",i,v)
	}

	// rune类型
	var n rune = '测'
	fmt.Printf("========%c\n",n)

	// 字符长度和字符串长度区别
	rune1  := []rune("我们不一样")
	rune2 := []rune("hahaha")
	fmt.Println(len(rune1),len(rune2))

	// []rune
	for i ,v := range []rune(str) {
		fmt.Printf("这是第三个循环---%d,%v\n",i,v)
	}

	// 更改字符串--直接更改是会报错的，需要转化为字符数组后再更改
	var byAr = []byte(str)
	fmt.Println(byAr)
	byAr[0]='s' // 注意这里是字符，不是字符串，'我'这样的中文是不可以的，会越界，需要rune
	// byAr[0]='我' 这种会越界，需要rune切片类型
	fmt.Println(byAr)


	// 再转回字符串
	str1 := string(byAr)
	fmt.Println(str1)
}