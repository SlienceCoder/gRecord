package main

import(
	"fmt"
	"strings"
	"gsty/L2/access"
)



func main(){

	// testString()
	// testbool()
	testInt()
}
func testInt(){
	var a int
	a = 8
	fmt.Println(a)

	a = -90
	fmt.Println(a)

	var b int
	b = 77889889778978
	fmt.Println(b)

	var c = 8.9
	fmt.Println(int(c))

}

func testbool(){
	a := true
	b := false
	if a == true && b ==true {
		fmt.Println("两个都是true")
	} else {
		fmt.Println("有一个不为true")
	}

	if a == true || b == true {
		fmt.Println("其中有一个为true")
	} else {
		fmt.Println("两个都不为true")
	}
}

func testString(){
	var a string
	a = "a:sfsdfd"
	fmt.Println(a)

	b := a
	fmt.Println("b----",b)

	c := "geuyfgdh\nfdsghfue"
	fmt.Println("c-----",c)

	// 万能占位符%v
	fmt.Printf("字符串占位符-%v  %v  %v",a,b,c)
	fmt.Printf("占位符--%s %s  %s",a,b,c)

	// ''字符串,原始打印，不进行转义
	d := `haha
	fds 
	\nffdgf
	dfdffds`
	fmt.Println("d ---",d)
	// len-求字符串的长度(所占的字节长度)
	d1 := "我们haha"
	d2 := "zhhaha"
	fmt.Println("字符串长度--",len(d1),len(d2))

	// + 拼接字符串
	fmt.Printf("%v\n",d1+d2)

	//Sprintf-格式化字符串
	e := fmt.Sprintf("%v%d","测试护具",8)
	fmt.Println(e)

	// Split-根据指定字符串分割，返回列表
	ips := "192.168.0.23;199.268.3.89"
	f := strings.Split(ips,";")
	fmt.Println(f)
	splitArr := "nihaoawoshishuinizhidaoma"
	f1 := strings.Split(splitArr,"a")
	fmt.Println(f1,len(f1))

	//Contains- 是否包含子串
	f2 := strings.Contains(splitArr,"wo")
	fmt.Println(f2)

	// HasPrefix
	f3 := strings.HasPrefix(splitArr,"nihao")
	fmt.Println(f3)
	// HasSuffix
	f4 := strings.HasSuffix(splitArr,"hah")
	fmt.Println(f4)

	// Index--正向，首次出现该子串的位置，没有该子串返回-1
	f5 := strings.Index(splitArr,"o00")
	fmt.Println(f5)

	// LastIndex --逆向，首次出现子串的位置，没有返回-1
	f6 := strings.LastIndex(splitArr,"o0")
	fmt.Println(f6)

	// Join -由一个切片根据子串拼接成字符串
	sli := []string{"--","==","~~"}
	f7 := strings.Join(sli,"ceshi")
	fmt.Println(f7)


}

func testAccess(){
	// 挎包访问外部变量
fmt.Println(access.A,access.B)

}