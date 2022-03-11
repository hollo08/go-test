package main

import (
    "fmt"
    "github.com/robfig/cron"
)

func main() {
    fmt.Println("start......")
    
    
    i := 0
    c := cron.New()
    spec := "*/5 * * * * ?"
    c.AddFunc(spec, func() {
        i++
        fmt.Printf("cron running: %d \n", i)
    })
    c.AddFunc("@every 1h1m", func() {
        i++
        fmt.Printf("cron running:%d\n", i)
    })
    c.Start()
    defer c.Stop()
    select {
    }
}
