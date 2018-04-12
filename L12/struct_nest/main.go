package main

import(
	"fmt"
)
type Address struct {
	a string
	b string
}

type User struct {
	name string
	age int
	addr *Address
}

func main(){
	u := &User{
		"kobe",89,
		&Address{
			"ceshi1","ceshi2",
		},

		
	}
	fmt.Printf("-----%#v\n",u)
}