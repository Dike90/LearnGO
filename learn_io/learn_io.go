package main

import (
	"fmt"
)

func main()  {
	var (
		name string
		age int
		//f float64
	)
	fmt.Println("请输入姓名:")
	fmt.Scanln(&name)

	fmt.Println("请输入年龄:")
	fmt.Scanf("%d",&age)
	//fmt.Scanf("%f",&f)
	fmt.Printf("hi,%s,you are %d years old!\n",name,age)
}