package main

import(
	"fmt"
)

// 结构体指针
type User struct{
	name string
	age int
	address string
}

func main(){
	// 1.结构体指针1
	var u1 *User = &User{
		"kobe",88,"losangles",
	}
	// 2.结构体指针2 结构体只声明是不可以赋值的
	var u2 *User = &User{}
	(*u2).name = "james"
	(*u2).age = 89
	(*u2).address = "beijing"

	// 结构体指针3
	var u3 *User = &User{
		name:"bosh",
		age:89,
		address:"shanghai",
	}

	// 4.结构体指针4
	var u4 *User = new(User)
	u4.name = "hah"
	u4.age = 90
	u4.address = "bfhjdsb"

	fmt.Println(u1,u2,u3,u4)

}