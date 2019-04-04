package main

import (
	"fmt"
)

func main() {
	fmt.Println("=================sliceLearn=================")
	sliceLearn()
	fmt.Println("=================mapLearn=================")
	mapLearn()
}

func sliceLearn() {
	sliceSrc := []int{1, 2, 3}
	sliceDst := make([]int, 8)
	res := copy(sliceDst, sliceSrc)
	fmt.Println(sliceDst, res)
	s1 := []int{4, 56, 8}
	s2 := []int{23, 546, 31}
	s1 = append(s1, 12, 34, 5) //append增加元素
	s2 = append(s2, s1...)     //追加另一个slice
	fmt.Println(s1)
	fmt.Println(s2)
}

func mapLearn() {
	var m1 map[string]int
	m1 = map[string]int{"k1": 124, "k2": 646}
	fmt.Println(m1)
	m2 := make(map[string]float64)
	m2["dk"] = 360000
	m2["dd"] = 12435
	m2["wk"] = 23451
	fmt.Println(m2)
}
