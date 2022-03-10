package main

import (
    "fmt"
    "github.com/koding/multiconfig"
)

type Server struct {
    Demo DemoConfig
}

type DemoConfig struct {
    Name    string
    Port    int
    Enabled bool
    Users   []string
}

func main() {
    fmt.Println("Hello, World!")
    m := multiconfig.NewWithPath("G:\\gowork\\src\\go-test\\test\\config.toml") // supports TOML, JSON and YAML
    
    serverConf := new(Server)
    err := m.Load(serverConf) // Check for error
    if err != nil {
        fmt.Println(err)
    }
    m.MustLoad(serverConf) // Panic's if there is any error
    
    fmt.Println(serverConf.Demo.Port)
    fmt.Println(serverConf.Demo.Name)
    fmt.Println(serverConf.Demo.Users)
}
