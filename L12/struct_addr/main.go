package main

import(
	"fmt"
)

// 内存分布
type test struct{
	A int32
	B int32
	C int32
	D int32
}

func main(){
	var t test
	fmt.Printf("--%p\n",&(t.A))
	fmt.Printf("--%p\n",&(t.B))
	fmt.Printf("--%p\n",&(t.C))
	fmt.Printf("--%p\n",&(t.D))

	/*
	结构体的内存布局是一块连续的空间
	--0xc420080010
--0xc420080014
--0xc420080018
--0xc42008001c
	*/
}