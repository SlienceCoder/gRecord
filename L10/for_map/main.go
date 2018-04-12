package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	var m map[string]int = make(map[string]int,1024)

	for i := 0;i < 20;i++{
		key := fmt.Sprintf("--%d",i)
		value := rand.Intn(1000)
		m[key]=value
	}

	for key ,value := range m {
		fmt.Printf("print %#v  %d\n",key,value)
	}
}