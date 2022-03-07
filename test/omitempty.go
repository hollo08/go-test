package main

import (
    "encoding/json"
    "fmt"
)

type Person2 struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Addr string `json:"addr,omitempty"`//omitempty:如果字段只为空(0, false, nil)，json就忽略这个字段
    Type int32 `json:"type,omitempty"`
}

func main() {
    p1 := Person2{
        Name: "taoge",
        Age: 30,
        Type: 0,
    }
    
    data, err := json.Marshal(p1)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("%s\n", data)
    fmt.Println(p1.Name, p1.Age, p1.Addr)
    
    p2 := Person2{
        Name: "Cang Laoshi",
        Age: 18,
        Addr: "Japan",
    }
    
    data2, err := json.Marshal(p2)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("%s\n", data2)
    
    fmt.Println(p2.Name, p2.Age, p2.Addr)
}
