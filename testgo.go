package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func say(s string) {
	for i := 0; i < 5; i++ {

		fmt.Println(s)
		runtime.Gosched()
	}
}

func main() {
	go say("world") //开一个新的Goroutines执行
	//say("hello") //当前Goroutines执行
	var chinese string = "汉字求长度"
	fmt.Println(len(chinese))
	var goos string = runtime.GOOS
	path := os.Getenv("PATH") //获取的path包含系统的和用户的
	pathArr := strings.Split(path, ";")
	fmt.Printf("the operation is %s\n", goos)
	for i, v := range pathArr {
		fmt.Printf("%d %s\n", i, v)
	}

}
