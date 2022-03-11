package main

import (
    "fmt"
    "sync"
)

func main() {
    //var wg sync.WaitGroup
    var once sync.Once
    onceBody := func() {
        fmt.Println("Only once")
    }
    done := make(chan bool)
    //var rev bool
    for i := 0; i < 10; i++ {
        go func() {
            once.Do(onceBody)
            done <- true
        }()
    }
    for i := 0; i < 3; i++ {
        //rev = <-done
        //fmt.Println(rev)
        select {
        case v, ok := <-done:
            if ok{
                fmt.Println(v)
            }else {
                fmt.Println("close")
            }
        default:
           fmt.Println("down")
        }
    }
}
