package main

import (
	"fmt"
)

func main() {
	greeting("Hello: ", "dk", "cxs", "sd")
	fmt.Println(min(1, 2, 378, 23))
	mySlice := []int{32, 45, 123}
	s := min(mySlice...)
	fmt.Println(s)
	deferLearn()
	fmt.Println("==============Test Defer===============")
	f1 := adder(3)
	fmt.Println("call adder:", f1(2))
	f2 := adder2()
	fmt.Println("call adder2", f2(4))
	closureTest(5, 80, 900)
	closureTest(1, 10, 200)
}
func greeting(prefix string, who ...string) {
	fmt.Printf(prefix)
	for i := 0; i < len(who); i++ {
		fmt.Printf(who[i])
		if i < len(who)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Println()

}

func min(n ...int) int {
	if len(n) == 0 {
		return 0
	}
	min := n[0]
	for _, v := range n {
		if v < min {
			min = v
		}
	}
	return min
}

func deferLearn() int {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d, ", i)
	}
	s := 42
	defer fmt.Println(s)
	return s
}

func adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func adder2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func closureAdder() func(int) int {
	var x int //定义一个局部变量
	return func(delta int) int {
		x += delta
		return x
	}
}

func closureTest(a, b, c int) {
	var f = closureAdder() //在f被销毁之前，闭包中的局部变量不会被销毁，导致closureAdder中x的值一直累积
	fmt.Print(f(a), "-")
	fmt.Print(f(b), "-")
	fmt.Println(f(c))
}
