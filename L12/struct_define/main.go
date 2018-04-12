package main

import(
	"fmt"
)

type User struct {
	name string
	age int
	address string
}

func main(){
	// 结构体初始化的三种方式
	// 1.
	var u1 = User{"kobe",12,"beijing"}
	// 2.
	var u2 User 
	u2.name = "james"
	u2.age = 90
	u2.address = "chigo"
	// 3.
	var u3 User = User{
		name:"wade",
		age:90,
		address:"fdsf",
	}

	fmt.Printf("--%#v\n",u1)
	fmt.Printf("--%#v\n",u2)
	fmt.Printf("--%#v\n",u3)
}