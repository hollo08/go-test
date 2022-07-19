package main

import (
	"fmt"
)
type TestStruct struct {
	name string
}
type TestInt32 int32

const (
	One TestInt32 = 1
	Two TestInt32 = 2
)

func main(){
	var t1 = "编辑推荐位，修改前/修改后（%s）"
	var t2 = "http://mp-api.dm.bolome.com/evidenceApply/%s/%s/t?%s"
	url1 := fmt.Sprintf(t1, "1", "2", "3")
	url2 := fmt.Sprintf(t2, "1", "2", "3")
	fmt.Println(url1)
	fmt.Println(url2)
	test := &TestStruct{}
	fmt.Println(test)
	fmt.Println(TestInt32(1))
	var t3 TestInt32
	t3 = TestInt32(3)
	fmt.Println(t3)
}
