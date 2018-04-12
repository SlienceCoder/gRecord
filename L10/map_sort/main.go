package main

import(
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 对map根据key进行排序

func main(){
	rand.Seed(time.Now().UnixNano())
	var m map[string]int = make(map[string]int,1024)

	for i := 0;i < 30;i++{
		key := fmt.Sprintf("stu%d",i)
		value := rand.Intn(1000)
		m[key]=value
	}


	// 创建数组，记录所有key
	var sortArr []string = make([]string,50)
	for key ,_ := range m{
		sortArr = append(sortArr,key)
	}

	sort.Strings(sortArr)

	for i,v := range sortArr{
		fmt.Printf("--%v %v %d\n",m[v],v,i)
	}

}