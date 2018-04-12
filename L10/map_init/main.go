package main

import(
	"fmt"
)

func main(){
	// 初始化1
	var d1 map[string]int=map[string]int{
		"1":1,
		"2":2,
		"3":56,
	}
	fmt.Printf("--%#v\n",d1)
	// 初始化方法2
	var d2 map[string]int
	d2 = make(map[string]int,12)
	d2["a"]=9
	d2["s"]=80

	fmt.Printf("--%#v\n",d2)

	// 通过一个key取值
	var key = "a"
	value ,ok := d2[key]
	if ok {
		fmt.Println(value)
	}
}