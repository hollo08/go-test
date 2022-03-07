package main

import "reflect"
import "fmt"

type Person3 struct {
    Name string
    Age  int
}

func main(){
    person3 := Person3{"zhangsan",30}
    v := reflect.ValueOf(&person3).Elem()
    fmt.Println(v)
    
    count := v.NumField()
    fmt.Println("count:",count)
    for i := 0; i < count; i++ {
        field := v.Field(i)
        fmt.Println(field)
        
        if field.Kind() == reflect.String {
            field.SetString("Test Value")
        }
    }
    fmt.Println(person3)
}

