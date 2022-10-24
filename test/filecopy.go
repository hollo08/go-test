package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

//编写拷贝文件的函数
//dstFileName：要拷贝到的地址
//srcFileName : 要拷贝的原文件
func CopyFile(dstFileName string,srcFileName string) (written int64,err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file err:",err)
	}
	//关闭流
	defer srcFile.Close()
	//获取到reader
	reader := bufio.NewReader(srcFile)
	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666) //0666 在windos下无效
	if err != nil {
		fmt.Println("open file err:",err)
		return
	}
	writer := bufio.NewWriter(dstFile)
	//关闭流
	defer dstFile.Close()
	//调用copy函数
	return io.Copy(writer,reader)
}

func getAllFiles(files []*fs.FileInfo, dir string){
	newFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range newFiles {
		if file.IsDir(){
			getAllFiles(files, dir + "\\" + file.Name())
			fmt.Println(file.Name())
		}else {
			files = append(files, &file)
			fmt.Println(len(files))
			srcFile := dir + "\\" + file.Name()
			dstFile := "C:\\迅雷下载\\小说包 密码123\\新建文件夹" + "\\" + file.Name()
			_, err := CopyFile(dstFile, srcFile)
			if err != nil {
				fmt.Println("文件拷贝错误",err)
			} else {
				fmt.Println("文件拷贝成功")
			}
		}
	}
}

func TestSlice(a *[]string){
	test := "tt"
	*a = append(*a, test)
}


func main() {
	//var aa []string
	dir := "C:\\迅雷下载\\小说包 密码123\\小说包 密码123"
	//aa = append(aa, dir)
	var files []*fs.FileInfo
	getAllFiles(files, dir)
	//TestSlice(&aa)
	//fmt.Println(len(aa))
	//var files []string
	//filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
	//	files = append(files, path)
	//	fmt.Println(info.Name())
	//	return nil
	//})
	//for _, f := range files{
	//	fmt.Println(f)
	//}
}