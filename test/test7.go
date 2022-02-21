package main

import (
	"fmt"
)

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main(){
	//testStruct1()
	book := testStruct2()
	fmt.Printf("book title is %s \n", book.title)
	testStruct3(book)
	fmt.Printf("book title is %s \n", book.title)
	testStruct4(&book)
	fmt.Printf("book title is %s \n", book.title)
	fmt.Printf("book title is %s \n", &book.title)
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