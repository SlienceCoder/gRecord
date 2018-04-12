package main

import(
	"fmt"
)

/*
	表明我们的map是引用逻辑
*/
func modifyArr(a map[string]int){
	a["123"]=123
}

func main(){
	var a map[string]int=map[string]int{
		"1":1,
		"2":2,
		"3":3,
	}

	b := a
	b["3"]=2344
	fmt.Printf("--%#v\n",a)
	modifyArr(b)
	fmt.Printf("--%#v\n",a)

}