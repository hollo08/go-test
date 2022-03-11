package main

import (
    "github.com/robfig/cron"
    "log"
)

type Hello struct {
    Str string
}

func(h Hello) Run() {
    log.Println(h.Str)
}

func main() {
    log.Println("Starting...")
    
    c := cron.New()
    h := Hello{"I Love You!"}
    // 添加定时任务
    c.AddJob("*/2 * * * * * ", h)
    // 添加定时任务
    c.AddFunc("*/5 * * * * * ", func() {
        log.Println("hello word")
    })
    
    s, err := cron.Parse("*/3 * * * * *")
    if err != nil {
        log.Println("Parse error")
    }
    h2 := Hello{"I Hate You!"}
    c.Schedule(s, h2)
    // 其中任务
    c.Start()
    // 关闭任务
    defer c.Stop()
    select {
    }
}
