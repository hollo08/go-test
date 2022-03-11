package main
import (
    "fmt"
    "time"
)
func main() {
    chann := make(chan int)
    go enqueue(chann)
    for {
        select {
        case v, ok := <-chann:
            if ok {
                fmt.Println(v)
            } else {
                fmt.Println("close")
                return
            }
        default:
            fmt.Println("waiting")
        }
    }
}
func enqueue(chann chan int) {
    time.Sleep(3 * time.Millisecond)
    chann <- 1
    close(chann)
}
