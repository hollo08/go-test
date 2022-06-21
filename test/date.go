package main

import (
    "fmt"
    "time"
)

func main(){
    now := time.Now()
    year, month, day := now.Date()
    afterDate := time.Date(year+3, month, day, 0, 0, 0,  0, time.Now().Location())
    beforeDate := time.Date(year, month, day-1, 0, 10, 0,  0, time.Now().Location())
    fmt.Println(now)
    fmt.Println(afterDate)
    fmt.Println(beforeDate.Format("2006-01-02 15:04:05"))
    fmt.Println(now.After(afterDate))
    fmt.Println(now.After(beforeDate))
}
