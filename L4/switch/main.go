package main

import(
	"fmt"
)
// 1.99乘法表
func ninenine(){
	for i:= 1;i <= 9;i++{
		for j := 1;j <= i;j++{
			fmt.Printf("%d*%d=%d\t",j,i,i*j)
		}
		fmt.Printf("\n")
	}
}

// 2.switch不带变量的多条件判断
func testSwitch1(){
	var num = 89
	switch {
	case num >=90 :
		fmt.Println("优秀")
	case num >80 && num <90:
		fmt.Println("良")
	case num >70&&num <= 80:
		fmt.Println("中")
	case num <=70&&num >=60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	
	}
}

// 3.fallthrough

func testSwitch2(){
	num := 88
	switch {
	case num >=90 :
		fmt.Println("优秀")
	case num >80 && num <90:
		fmt.Println("良")
		fallthrough
	case num >70&&num <= 80:
		fmt.Println("中")
	case num <=70&&num >=60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}

// 4.switch带赋值和变量，合并多个case

func testSwitch3(){
	num := 99

	switch  num{
	case 1,2,3,4,5,99:
		fmt.Println("bingo")
	case 88,66:
		fmt.Println("双击666")
	default:
		fmt.Println("default")
	}
}

// 5.调用函数获取值
func getValue()int{
	return 8
}
func testSwitch4(){
	switch num:=getValue();num{
	case 1,2,3,4,5,99:
		fmt.Println("bingo")
	case 88,66:
		fmt.Println("双击666")
	default:
		fmt.Println("default")
	
	}

}

// 6.普通的switch
func testSwitch5(){
	num := 90
	switch num {
	case 1:
		fmt.Println("haha")
	case 90:
		fmt.Println("无聊")
	case 77:
		fmt.Println("77")
	default:
		fmt.Println("哎呀--default啦")
	}
}

func main(){
// ninenine()
// testSwitch()
// testSwitch3()
testSwitch5()
}