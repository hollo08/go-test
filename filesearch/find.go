package main

import (
	"fmt"
	f "github.com/hollo08/go-test/file"
)

func main(){
	f.GetDb()
	name := "陈素怡"
	result := make(map[string]string, 0)
	ff := make([]f.Files, 0)
	//f.Db.Debug().First(ff, 1)
	//f.Db.Debug().Last(ff)
	f.Db.Debug().Where("path like (?)", "%"+name+"%").Find(&ff)
	//f.Db.Debug().Where("name like (?)", "%"+name+"%").Find(&ff)
	//f.Db.Debug().Where("file_type = (?)", name).Find(&ff)
	for _, fff := range ff{
		result[fff.Path] = fff.Name
	}
	for k, v := range result{
		fmt.Println(fmt.Sprintf("%s,%s", k, v))
	}
}

