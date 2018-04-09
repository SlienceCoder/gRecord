package main

import (
	"fmt"
	"time"
)

func test(){
	for i := 0; i < 10; i++ {
		fmt.Printf("index --%d\n",i)
	}
	fmt.Println("end--")
}

func main (){
	go test()
	
	time.Sleep(time.Second*5)
	fmt.Println("all_end")
}