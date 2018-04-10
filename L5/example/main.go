package main

import(
	"fmt"
)

// 1. 求1到100之内的所有质数，并打印到屏幕上
func judgeZhishu(n int)bool{
	if n <=1 {
		return false
	}
	for i := 2;i < n;i++ {
		if n%i==0{
			return false
		}

	}
	fmt.Printf("我是质数----%d\n",n)
	return true
}

func testZhishu(){
	for i :=1;i <= 100;i++{
		judgeZhishu(i)
	}
}

/*
	2. 打印出所有的“ 仙花数”，所谓“ 仙花数”是指 个三位数，其各位数字  和 等于该数本身。 如:153是 个“ 仙花数”，因为153=1的三次 +5的三次  +3的三次 。求100到1000之间的所有 仙花数，
*/
func judgeShuixianhua(n int)bool{
	a := n%10
	b := (n/10)%10
	c := (n/100)

	if n == a*a*a+b*b*b+c*c*c {
		fmt.Printf("我是水仙花--%d\n",n)
		return true
	}
	return false
}
func testShuixianhua(){
	for i:=100 ;i <1000;i++{
		judgeShuixianhua(i)
	}
}


// 3. 输   字符，分别统计出其中英 字 、空格、数字和其它字符的个数

func testStr(s string)(alpCount int,spaCount int,numCount int,otherCount int){

	var sArr []rune = []rune(s)

	for _ ,v := range sArr{
		fmt.Println(v)
		if (v>='a'&&v<='z')||(v>='A'&&v<='Z'){
			alpCount += 1
		}else if v == ' '{
			spaCount += 1
		} else if v>='0'&&v<='9' {
			numCount += 1
		} else {
			otherCount += 1
		}

	}
	return

}

func main(){
	// testZhishu()
	// fmt.Println(judgeZhishu(7))
	// testShuixianhua()
	a,b,c,d := testStr("aodiu  42期,'9")
	fmt.Println(a,b,c,d)
}