package main

import (
	"fmt"
	"reflect"
)

type myInt int

type unkownType struct {
	s1 ,s2 ,s3 string
}

func (t unkownType) String() string  {
	return t.s1 +"-"+t.s2+"-"+t.s3
}

func (t unkownType) Upper(str string) string {
	return t.s1 +str +t.s2 +str +t.s3
}

func main() {
	var x int64 = 42
	fmt.Println(reflect.TypeOf(x), reflect.ValueOf(x))
	var myint myInt = 45
	fmt.Println(reflect.TypeOf(myint), reflect.ValueOf(myint))
	v := reflect.ValueOf(myint) //v的类型为reflect.Value
	fmt.Printf("the type of v is :%T\n",v) //the type of v is :reflect.Value
	fmt.Println(v.Kind()) //int Kind会返回底层值的类型
	fmt.Println(v.Interface()) //45
	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println(v.Type())
	fmt.Println("settability of v:", v.CanSet())
	v = v.Elem()
	fmt.Println("settability of v:", v.CanSet())
	v.SetInt(36)
	fmt.Println(x)
	var secret interface{} = unkownType{"abc","wed","cat"}
	val := reflect.ValueOf(secret)
	typ := reflect.TypeOf(secret)
	fmt.Println(typ)
	knd := val.Kind()
	fmt.Println(knd) //struct
	for i:=0; i< val.NumField() ; i++{  //通过NumField来获取结构体字段的数量
		fmt.Printf("Field %d:%v\n", i , val.Field(i)) //通过Field来获取字段的值
	}

	fmt.Println("unkownType的方法数量为",val.NumMethod()) //通过NumMethod来获取结构体方法的数量
	res1 := val.Method(0).Call(nil)
	//带参数的调用
	params := make([]reflect.Value,1)
	str := "--"
	params[0] = reflect.ValueOf(str)
	res2:=val.Method(1).Call(params)
	fmt.Println(res1)
	fmt.Println(res2)

}
