package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

type TagType struct {
	field1 bool   "This is a boolean"
	field2 string "This is a string"
	field3 int    "This is a int"
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Println(ixField.Type)
	fmt.Println(ixField.Name)
	fmt.Printf("%v\n", ixField.Tag)
}

//struct匿名字段与组合
type innerS struct {
	in1 int
	in2 int
}
type outerS struct {
	a string
	b string
	int
	innerS
}

func main() {
	var p1 Person
	p1.firstName = "chirst"
	p1.lastName = "underwood"
	upPerson(&p1)
	fmt.Println(p1)

	p2 := new(Person) //new返回一个*Person类型的值
	p2.firstName = "jack"
	p2.lastName = "underwood"
	upPerson(p2)
	fmt.Println(p2.firstName, p2.lastName)
	p3 := &Person{"Mr.", "Stepool"} //与new是等价的
	upPerson(p3)
	fmt.Println(p3.firstName, p3.lastName)

	tt := TagType{true, "Ms.Clare", 32}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}

	outer := new(outerS)
	outer.a = "this is outer a"
	outer.int = 42
	outer.in1 = 12
	fmt.Println("this is innerS.in1:", outer.in1)
	fmt.Println("anonymous field is ", outer.int)
}
