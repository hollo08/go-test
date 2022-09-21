package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var CHARS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

/*RandAllString  生成随机字符串([a~zA~Z0~9])
  lenNum 长度
*/
func RandAllString(lenNum int) string {
	str := strings.Builder{}
	length := len(CHARS)
	time.Sleep(1)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < lenNum; i++ {
		l := CHARS[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

/*RandNumString  生成随机数字字符串([0~9])
  lenNum 长度
*/
func RandNumString(lenNum int) string {
	str := strings.Builder{}
	length := 10
	for i := 0; i < lenNum; i++ {
		str.WriteString(CHARS[52+rand.Intn(length)])
	}
	return str.String()
}

/*RandString  生成随机字符串(a~zA~Z])
  lenNum 长度
*/
func RandString(lenNum int) string {
	str := strings.Builder{}
	length := 52
	for i := 0; i < lenNum; i++ {
		str.WriteString(CHARS[rand.Intn(length)])
	}
	return str.String()
}

const GoTime = "2006-01-02 15:04:05.999999"

func main() {
	for i := 0; i < 10000000; i++ {
		randStr := RandString(10)
		randNumStr := RandNumString(10)
		randAllStr := RandAllString(10)
		fmt.Println(i, time.Now().Format(GoTime), randStr, randNumStr, randAllStr)
	}
}
