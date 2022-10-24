package main

import (
	"fmt"
	f "github.com/hollo08/go-test/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"path"
	"strings"
	"time"
)

type lastUpdate struct {
	gorm.Model
}

var pathSeparator = string(os.PathSeparator)
func traverseDir(filePath string) error{
	file, err := os.Open(filePath)
	if err != nil {
		if strings.Contains(err.Error(), "Access is denied"){
			return nil
		}
		return err
	}
	fileInfo, err := file.Readdir(0)
	if err != nil {
		return err
	}

	for _, value := range fileInfo {
		if value.IsDir() {
			//dirNames = append(dirNames, value.Name())
			//fmt.Println(value.Name())
			err = traverseDir(filePath+pathSeparator+value.Name())
			if err != nil {
				return err
			}
		}else{
			fmt.Println(filePath+pathSeparator+value.Name())
			writeToDb(value.Name(), filePath, getExt(value.Name()), value.ModTime())
		}
	}
	return err
}
func main() {
	f.GetDb()
	var filePath = "Z:\\"
	f.Db.SingularTable(true)
	f.Db.AutoMigrate(&f.Files{})
	err := traverseDir(filePath)
	if err != nil {
		fmt.Println(err)
	}
	f.Db.Close()
}
func getExt(name string) string{
	ext := path.Ext(name)
	if ext != ""{
		ext = ext[1:]
	}
	return ext
}
func writeToDb(name, path, fileType string, lastMod time.Time){
	//自动检查 file 结构是否变化，变化则进行迁移
	//db.AutoMigrate(&file{})
	f.Db.Create(&f.Files{
		Name:     name,
		Path:     path,
		FileType: fileType,
		DirModTime: lastMod,
	})
}

