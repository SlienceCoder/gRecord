package main

import(
	"fmt"
	"time"
)

func main(){
	// testTime()
	// testTIme2(1523261913)
	// testTimeFormat()
	// 定时器
	// testTick()
	// testAlltimeSec()

}

func testTick(){
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Printf("%v\n",i)
		
	}
}
func testAlltimeSec(){
	// 测试一段代码执行时间
	start := time.Now().UnixNano()
	time.Sleep(time.Second*5)
	end := time.Now().UnixNano()

	temp := (end-start)/1000/1000/1000

	fmt.Println(temp)
	fmt.Printf("%d\n",time.Second)
	fmt.Printf("%d\n",time.Millisecond)
	fmt.Printf("%d\n",time.Microsecond)
	fmt.Printf("%d\n",time.Nanosecond)

}

func testTimeFormat(){
	// 格式化方法1
	now := time.Now()
	timeStr := now.Format("2006/01/06 15:04:05")
	fmt.Printf("time :  %s\n",timeStr)

	// 格式化方法2
	s1 := fmt.Sprintf("林外一个:%02d/%02d/%02d %02d:%02d:%02d",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Println(s1)
}


func testTIme2(temp int64){
	// 根据给定的时间戳求时间
	t := time.Unix(temp,0)

	year := t.Year()
	month := t.Month()
	day := t.Day()
	hour := t.Hour()
	min := t.Minute()
	sec := t.Second()

	fmt.Printf("\n%02d-%02d-%02d %02d:%02d:%02d\n\n",year,month,day,hour,min,sec)


}
// 年月日时间，以及时间戳
func testTime(){
	now := time.Now()
	// fmt.Println(now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	min := now.Minute()
	sec := now.Second()

	fmt.Printf("\n%02d-%02d-%02d %02d:%02d:%02d\n\n",year,month,day,hour,min,sec)

	fmt.Println("时间戳--",now.Unix())
	/*
	2018-04-09 16:18:33

时间戳-- 1523261913
	*/
}