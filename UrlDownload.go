package main

import (
    "bufio"
    "fmt"
    "io"
    "net/http"
    "os"
    "path"
)
// 文件下载到本地，通常的思路就是先获得网络文件的 输入流 以及本地文件的 输出流 ，再读取输入流到输出流中，因此自然也要获取相应的 Reader 和 Writer 。

func main() {
    imgPath := `F:\order\`
    imgUrl := "https://bolome-file.oss-cn-hangzhou.aliyuncs.com/ORDER_TRADING_20211005_eb15a1821c414d259f01b167b1effc7e.zip"
    
    fileName := path.Base(imgUrl)
    
    
    res, err := http.Get(imgUrl)
    if err != nil {
        fmt.Println("A error occurred!")
        return
    }
    defer res.Body.Close()
    // 获得get请求响应的reader对象
    reader := bufio.NewReaderSize(res.Body, 32 * 1024)
    
    
    file, err := os.Create(imgPath + fileName)
    if err != nil {
        panic(err)
    }
    // 获得文件的writer对象
    writer := bufio.NewWriter(file)
    
    written, _ := io.Copy(writer, reader)
    fmt.Printf("Total length: %d", written)
}

