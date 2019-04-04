package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x int64 = 42
	fmt.Println(reflect.TypeOf(x), reflect.ValueOf(x))
}
