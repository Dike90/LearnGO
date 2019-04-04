package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
)

type twoInts struct {
	int1 int
	int2 int
}

type Shaper interface {
	Area() float64
}

type rectangle struct {
	length, width float64
}

func (r *rectangle) Area() float64 {
	return r.length * r.width
}

type circle struct {
	rad float64
}

func (c circle) Area() float64 {
	return math.Pi * c.rad * c.rad
}

//空接口变量可以被赋予任何值，所有的类型都实现了空接口
type Any interface{}

type person struct {
	name string
	age  int
}

func main() {
	twoNumber := &twoInts{12, 42}
	fmt.Println(twoNumber.AddThem())
	twoNumber.Modify(1, 2)
	fmt.Println(twoNumber)
	twoNumber.Modify2(3, 4)
	fmt.Println(twoNumber)
	//fmt.Println(twoNumber)
	//获取分配内存的大小
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Allocate %d Kb\n", m.Alloc/1024)
	/************************interface和多态**************************/
	r := rectangle{3, 4}
	c := circle{6}            //
	shapes := []Shaper{&r, c} //若receiver为指针类型,传递给接口变量时，需要取地址，否则编译报错
	for n, _ := range shapes {
		fmt.Println("shape detail:", shapes[n])
		fmt.Println("Area of this shape is :", shapes[n].Area())
	}
	var shape Shaper
	shape = &r
	//可以使用如下方法判断接口变量保存的实例变量的实际类型
	if v, ok := shape.(*rectangle); ok {
		fmt.Printf("the type of shape is :%T\n", v)
	}
	if v, ok := shape.(circle); ok {
		fmt.Printf("the type of shape is :%T\n", v)
	} else {
		fmt.Println("the shape is not a circle")
	}
	/************************空接口**************************/
	var any Any
	var i = 5
	any = i
	fmt.Printf("the type is:%T,the value is:%v\n", any, any)
	var str = "asdfj"
	any = str
	fmt.Printf("the type is:%T,the value is:%v\n", any, any)
	pers := &person{"dk", 28}
	any = pers
	fmt.Printf("the type is:%T,the value is:%v\n", any, any)
	switch t := any.(type) { //any.(type)必须在type-switch结构中使用，不能单独使用
	case string:
		fmt.Printf("this is a %T,and value is:%v\n", t, t)
	case int:
		fmt.Printf("this is a %T,and value is:%v\n", t, t)
	case person:
		fmt.Printf("this is a %T,and value is:%v\n", t, t)
	case *person:
		fmt.Printf("this is a %T,and value is:%v\n", t, t) //this is a *main.person,and value is:&{dk,28}
	}
}

func (t *twoInts) AddThem() int {
	return t.int1 + t.int2
}

func (this *twoInts) Modify(int1, int2 int) {
	this.int1 = int1
	this.int2 = int2
}

//非指针的接收者无法更改原实例的值
func (this twoInts) Modify2(int1, int2 int) {
	this.int1 = int1
	this.int2 = int2
}

//string方法是go中的特殊方法，fmt.Print族函数会自动调用
func (t *twoInts) String() string {
	return "the first int is:" + strconv.Itoa(t.int1) + ",the second int is:" + strconv.Itoa(t.int2)
}
