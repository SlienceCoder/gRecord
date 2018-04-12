package main

import(
	"fmt"
)

func main(){
	var m map[int]int = map[int]int{1:1,2:2,3:3,4:4}
	fmt.Printf("---%#v\n",m)
	for i,_ := range m{
		delete(m,i)
	}
	fmt.Printf("--%#v\n",m)
}