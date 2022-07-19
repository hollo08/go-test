package main

import "fmt"

func main(){
    var weakDay string
    fmt.Scan(&weakDay)
    switch weakDay {
    case "Mon.":
        fmt.Println(1)
    case "Tues.":
        fmt.Println(2)
    case "Wed.":
        fmt.Println(3)
    case "Thurs.":
        fmt.Println(4)
    case "Fri.":
        fmt.Println(5)
    case "Sat.":
        fmt.Println(6)
    case "Sun.":
        fmt.Println(7)
    default:
        panic("wrong day")
    }
}
