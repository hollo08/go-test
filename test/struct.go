package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main(){
	book1 := &Books{}
	fmt.Println(book1.title == "")
	//testStruct1()
	book := testStruct2()
	fmt.Printf("book title is %s \n", book.title)
	testStruct3(book)
	fmt.Printf("book title is %s \n", book.title)
	testStruct4(&book)
	fmt.Printf("book title is %s \n", book.title)
	fmt.Printf("book title is %s \n", &book.title)
	student:=student{"小明",18, time.Now().Unix()}
	if result,err:=json.Marshal(&student);err==nil{
		fmt.Println(string(result))
	}
	teacher:=teacher{"小明",18, time.Now().Unix()}
	if result,err:=json.Marshal(&teacher);err==nil{
		fmt.Println(string(result))
	}
}

func testStruct1(){
	var book Books
	book = Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	fmt.Printf("book authis is : %s and book is is : %d \n", book.author, book.book_id)
}

func testStruct2() Books{
	var book Books
	book.title = "Go Language"
	book.author = "google"
	book.subject = "about go language"
	book.book_id = 2122
	//fmt.Printf("book authis is : %s and book is is : %d \n", book.author, book.book_id)
	return book
}

func testStruct3(book Books){
	book.title = "Change1"
}
func testStruct4(book *Books){
	book.title = "Change2"
}

type student  struct{
	Name  string   `json:"namesss"`//标记json名字为name　　　
	age    int     `json:"age"`//小写开头的，是private属性
	Time int64    `json:"-"`        // 标记忽略该字段
}

type teacher  struct{
	Name  string   `form:"namesss"`//标记json名字为name　　　
	Age    int     `form:"age"`
	Time int64    `form:"-"`        // 标记忽略该字段
}

