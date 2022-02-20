package main

import "fmt"

type Person interface {
	GetAge() int
	GetName() string
	SetAge(age int)
	SetName(name string)
}

type Car interface {
	GetAge() int
	GetName() string
}

type Student struct {
	age  int
	name string
}

func(s *Student) GetAge() int {
	return s.age
}

func(s *Student) GetName() string {
	return s.name
}

func(s *Student) SetAge(age int){
	s.age = age
}

func(s *Student) SetName(name string){
	s.name = name
}

func main() {
	var p1 Person
	p1 = new(Student)
	p1.SetAge(1)
	p1.SetName("peter")
	fmt.Println("This person name is", p1.GetName())

	var p Person = &Student{20, "Elon"}

	fmt.Println("This person name is", p.GetName())

	fmt.Println("This person age is", p.GetAge())

	var c Car = &Student{1, "BMW"}

	fmt.Println("This car name is", c.GetName())

	fmt.Println("This car age is", c.GetAge())

	switch t := p.(type) {
		case *Student:
			fmt.Printf("student %T \n", t)
	}
	if  t,ok := c.(*Student); ok{
		fmt.Printf("t implements Student %T", t)
	}
}
