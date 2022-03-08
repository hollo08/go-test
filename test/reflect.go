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
    reflect2()
}

func reflect2(){
    var x float64 = 5.3
    fmt.Println("type:", reflect.TypeOf(x))
    v := reflect.ValueOf(x)
    fmt.Println("value:", v)
    fmt.Println("type:", v.Type())
    fmt.Println("kind:", v.Kind())
    fmt.Println("value:", v.Float())
    fmt.Println(v.Interface())
    fmt.Printf("value is %5.2e\n", v.Interface())
    y := v.Interface().(float64)
    fmt.Println(y)
}

