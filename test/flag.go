package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	cfg := flag.String("c", "cfg.dev.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	commonCfg := flag.String("cc", "cfg.common.json", "common configuration file")
	flag.Parse()
	fmt.Println(*cfg)
	fmt.Println(*version)
	fmt.Println(*commonCfg)
	base := time.Now().Format("060102")
	timestamp := time.Now().UnixNano() / 1e8
	fmt.Println(base)
	fmt.Println(timestamp)
	a := strconv.FormatInt(timestamp, 10)
	fmt.Println(a[len(a)-6:])
	fmt.Println(fmt.Sprintf("%s%s", base, a[len(a)-6:]))
	fmt.Println(GenCharShowId())
	fmt.Println(time.Now().UnixNano())
	time.Sleep(10)
	fmt.Println(time.Now().UnixNano())
}

func GenCharShowId() string {
	result := make([]byte, 0)
	var base int64 = 13
	timestamp := fmt.Sprintf("%10v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000000))
	fmt.Println(rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000000))
	num, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return ""
	}
	for num > 0 {
		result = append(result, showIdChars[num%base])
		num = num / base
	}
	return string(result)
}

var showIdChars = [36]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P',
	'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
