package main

import(
	"fmt"
)
const(
	a = 999
	b = iota
	c = iota
	d = iota
	e = 8
	f = 7
	g = iota
)
const(
	h = iota
	i 
)
func main(){
fmt.Println(a,b,c,d,e,f,g,h,i)
}