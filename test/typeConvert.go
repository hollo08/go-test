package main

import (
	"fmt"
	"strconv"
)

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}


func Factorial(n uint64)(result uint64) {
	if (n > 0) {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func testTypeConvert(){
	var i int = 12
	var f float32 = 12.11
	var s string = "ss"
	var ss string = "3"
	fmt.Printf("i's value is %d, f's value is %f, s's value is %s \n", i, f, s)
	i = 63234234234234
	fmt.Printf("i/f = %d, i/f=%f \n", i/10, float32(i))
	//sss, err := strconv.Atoi(ss)
	sss, err := strconv.ParseInt(ss, 10, 64)
	if(err != nil){
		fmt.Printf(err.Error())
	}else {
		fmt.Printf("convert ss to int, ss's value is %d \n", uint64(sss))
	}

}

func testConvert(){
	//转正整数
	str1 := "123"
	int1, err := strconv.Atoi(str1)
	if(err != nil){
		fmt.Printf(err.Error())
	}else{
		fmt.Printf("int1's value is %d, type is %T \n", int1, int1)
	}
	//转负整数
	str2 := "-123"
	int2, err := strconv.Atoi(str2)
	if(err != nil){
		fmt.Printf(err.Error())
	}else{
		fmt.Printf("int2's value is %d, type is %T \n", int2, int2)
	}
	//转int64
	str3 := "123"
	int3, err := strconv.ParseInt(str3, 10, 64)
	if(err != nil){
		fmt.Printf(err.Error())
	}else{
		fmt.Printf("int3's value is %d, type is %T \n", int3, int3)
	}
	//转 8进制int64
	str4 := "123"
	//ParseInt 三个参数，1：要转换的字符串，2：要转换的进制包含2,8,10,16，3：要转换的数字类型int8(8), int16(16), int32(32), int64(64)
	int4, err := strconv.ParseInt(str4, 8, 64)
	if(err != nil){
		fmt.Printf(err.Error())
	}else{
		fmt.Printf("int4's value is %d, type is %T \n", int4, int4)
	}
	//转8进制int32
	str5 := "123"
	//ParseInt 三个参数，1：要转换的字符串，2：要转换的进制包含2,8,10,16，3：要转换的数字类型int8(8), int16(16), int32(32), int64(64)
	int5, err := strconv.ParseInt(str5, 10, 32)
	if(err != nil){
		fmt.Printf(err.Error())
	}else{
		fmt.Printf("int5's value is %d, type is %T \n", int5, int5)
	}
	//转包含其它字符的数字，不能转换
	str6 := "123abc"
	//ParseInt 三个参数，1：要转换的字符串，2：要转换的进制包含2,8,10,16，3：要转换的数字类型int8(8), int16(16), int32(32), int64(64)
	int6, err := strconv.ParseInt(str6, 10, 64)
	if(err != nil){
		fmt.Printf(err.Error())
		fmt.Printf("\n")
	}else{
		fmt.Printf("int6's value is %d, type is %T \n", int6, int6)
	}
	//转超大的正整数，转换失败
	str7 := "111111111111111111111111111"
	int7, err := strconv.Atoi(str7)
	if(err != nil){
		fmt.Printf(err.Error())
		fmt.Printf("\n")
	}else{
		fmt.Printf("int7's value is %d, type is %T \n", int7, int7)
	}
	//转浮点数字符串,转换失败，
	str8 := "11.11"
	int8, err := strconv.Atoi(str8)
	if(err != nil){
		fmt.Printf(err.Error())
		fmt.Printf("\n")
	}else{
		fmt.Printf("int8's value is %d, type is %T \n", int8, int8)
	}
	//转浮点数字符串,转换失败，
	str9 := "11.11"
	int9, err := strconv.ParseFloat(str9, 32)
	if(err != nil){
		fmt.Printf(err.Error())
		fmt.Printf("\n")
	}else{
		fmt.Printf("int9's value is %f, type is %T \n", int9, int9)
	}
}

func testNumberConvert(){
	var a int64 = 12121212
	var b int = 111
	fmt.Println(strconv.Itoa(b) + "dddd")
	fmt.Println(strconv.FormatInt(a, 10))
	//string到int

	var c string = "123"
	i, _ :=strconv.Atoi(c)
	fmt.Println(i)
	//string到int64

	i64, _ := strconv.ParseInt(c, 10, 64)
	fmt.Println(i64)

	//int到string
	s:=strconv.Itoa(i)
	fmt.Println(s)

	//int64到string
	s64 := strconv.FormatInt(i64,10)
	fmt.Println(s64)

	//string到float32(float64)
	f,_ := strconv.ParseFloat(s,32)
	fmt.Println(f)
	f64,_ := strconv.ParseFloat(s,64)
	fmt.Println(f64)

	//float到string
	sf := strconv.FormatFloat(f, 'E', -1, 32)
	fmt.Println(sf)

	sf64 := strconv.FormatFloat(f, 'E', -1, 64)
	fmt.Println(sf64)

	// 'b' (-ddddp±ddd，二进制指数)

	// 'e' (-d.dddde±dd，十进制指数)

	// 'E' (-d.ddddE±dd，十进制指数)

	// 'f' (-ddd.dddd，没有指数)

	// 'g' ('e':大指数，'f':其它情况)

	// 'G' ('E':大指数，'f':其它情况)
}

func main() {
	var i int
	for i = 1; i <= 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
	fmt.Printf("\n")
	i = 10
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
	testTypeConvert()
	testConvert()
	testNumberConvert()
}
