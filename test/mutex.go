package main

import (
    "fmt"
    "sync"
)

func main() {
    //concurrentProblem()
    //lockMap()
    syncMap()
}

func concurrentProblem(){
    mp := make(map[string]string)
    go func() {
        for {
            mp["hello"] = "world"
        }
    }()
    go func() {
        for {
            _ = mp["hello"]
        }
    }()
    select {}
}

func syncMap(){
    mp := sync.Map{}
    go func() {
        for {
            mp.Store("hello", "world")
        }
    }()
    go func() {
        for {
            t, _ := mp.Load("hello")
            fmt.Println(t)
        }
    }()
    select {}
}

func lockMap(){
    mp := MyMap{
        RWMutex: sync.RWMutex{},
        mp:      make(map[string]string),
    }
    go func() {
        for {
            mp.Set("hello", "world")
        }
    }()
    go func() {
        for {
            t, _ := mp.Get("hello")
            fmt.Println(t)
        }
    }()
    select {}
}

type MyMap struct {
    sync.RWMutex
    mp map[string]string
}

func (m *MyMap) Get(key string) (value string, ok bool) {
    m.RLock()
    defer m.RUnlock()
    value, ok = m.mp[key]
    return
}

func (m *MyMap) Set(key, value string) {
    m.Lock()
    defer m.Unlock()
    m.mp[key] = value
}