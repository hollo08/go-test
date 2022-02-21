package main

import "fmt"

func main(){
	//testPoint1()
	//testPoint2()
	//testPoint3()
	//var a, b int
	a, b := 3, 4
	testPoint4(&a, &b)
	fmt.Printf("a's value is : %d \n", a)
	fmt.Printf("b's value is : %d \n", b)
}

func testPoint4(a, b *int){
	//var temp int
	//temp = *a
	//*a = *b
	//*b = temp
	*a, *b = *b, *a
}

func testPoint1(){
	var a int = 1
	var i *int
	i = &a
	var nullI *int
	fmt.Printf("a's address is %x \n", &a)
	fmt.Printf("i's address is %x \n", i)
	fmt.Printf("*i is %d \n", *i)
	if(nullI != nil){
		fmt.Printf("nullI's value is %x \n", *nullI)
	}else {
		fmt.Printf("*nullI is nil and wasn't assigned address \n")
	}
	//
	fmt.Printf("*nullI address is %d \n", nullI)
}

func testPoint2(){
	a := []int{10,100,200}
	var i int
	var ptr [3]*int;

	for  i = 0; i < 3; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for  i = 0; i < 3; i++ {
		fmt.Printf("ptr[%d]'s address is %x\n", i, ptr[i] )
	}

	for i = 0; i < 3; i++ {
		fmt.Printf("ptr[%d]'s value is %d\n", i, *ptr[i])
	}

	for i = 0; i < 3; i++ {
		fmt.Printf("a[%d]'s address is %x\n", i, &a[i])
	}
}

func testPoint3(){
	var a int
	var ptr *int
	var pptr **int
	var ppptr **(*int)
	var ppptr2 ***int
	a = 3000
	/* 指针 ptr 地址 */
	ptr = &a
	/* 指向指针 ptr 地址 */
	pptr = &ptr
	/*指向指向指针 pptr 地址*/
	ppptr = &pptr
	ppptr2 = &pptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a )
	fmt.Printf("指针变量 *ptr = %d\n", *ptr )
	fmt.Printf("指针变量 ptr 的地址 = %x\n", ptr )
	fmt.Printf("指针变量 &a 的地址 = %x\n", &a )
	fmt.Printf("指针变量 pptr 的地址 = %x\n", pptr )
	fmt.Printf("指针变量 &ptr 的地址 = %x\n", &ptr )
	fmt.Printf("指针变量 *ppptr2 的地址 = %x\n", *ppptr2)//等效于&**ppptr2
	fmt.Printf("指针变量 &pptr 的地址 = %x\n", &pptr )
	fmt.Printf("指针变量 ppptr 的地址 = %x\n", ppptr )
	fmt.Printf("指针变量 ppptr2 的地址 = %x\n", ppptr2 )
	fmt.Printf("指针变量 &ppptr 的地址 = %x\n", &ppptr )
	fmt.Printf("指针变量 &ppptr2 的地址 = %x\n", &ppptr2 )
	fmt.Printf("指针变量 **ppptr2 的地址 = %x\n", **ppptr2)//等效于&***ppptr2
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
	fmt.Printf("指向指针的指针变量 **(*ppptr) = %d\n", **(*ppptr))
	fmt.Printf("指向指针的指针变量 ***ppptr2 = %d\n", ***ppptr2)
}

