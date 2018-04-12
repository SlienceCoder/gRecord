package main

import(
	"fmt"
)

// 切片里面装map
func testSliceMap(){
	var s []map[string]int=make([]map[string]int,5,10)

	for i,v := range s {
		fmt.Printf("-------%d  %v\n",i,v)
	}

	s[0]= make(map[string]int,8)
	s[0]["0"]=11
	s[0]["1"]=90
	s[0]["2"]=67
	for i,v := range s{
		fmt.Printf("-----%d  %v\n",i,v)
	}


}

func testMapSlice(){
	var m map[string][]int 
	m = make(map[string][]int,12)
	// m["1"]=make([]int,5,10)
	// m["2"]=make([]int,3,8)
	_ ,ok := m["1"]
	if !ok {
		m["1"]=make([]int,5,10)
	}
	m["1"] = append(m["1"],200)
	m["1"] = append(m["1"],300)
	m["1"] = append(m["1"],400)
	fmt.Printf("----%#v\n",m)


}

func main(){
// testSliceMap()
testMapSlice()
}