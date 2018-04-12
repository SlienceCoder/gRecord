package main

import(
	"fmt"
)
func main(){
	var d map[string]int

	fmt.Println(d)

	if d == nil {
		d = make(map[string]int,10)
		d["1"]=1
		d["2"]=2
		d["3"]=3
		fmt.Printf("---%#v\n",d)
	}

}