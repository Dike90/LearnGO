package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unsafe"
)

var x int //自动初始化为0
var f float32
var s = "abc" //自动推断类型
//定义多个变量
var y, z, w int //相同类型的多个变量
//以组的方式整理多行变量
var (
	a    int
	b    bool
	c, d = 100, "cde"
)

//常量的值必须是编译期可确定的字符、字符串、数字或布尔值
//常量的类型可由编译器推断得出
const x1, y1 int = 124, 0x42
const s1 = "hello, world"
const c1 = '我'
const (
	i1, f1 = 12, 0.123
	b1     = true
)

//常量也可以是某些编译器能计算出结果的表达式
//如 unsafe.Sizeof len cap等
const (
	ptrSize = unsafe.Sizeof(uintptr(0))
	strSize = len("你好!")
)
const (
	x2 uint16 = 120
	y2        //类型为uint16,值为 120
	s2 = "abc"
	z2 //类型为string，值为 abc
)

//t := "error" //函数外面不能使用:=
func main() {
	fmt.Println(&x, x)
	//一个新的变量，并自动推断类型，会覆盖全局变量x
	x := 123                       //不能提供类型，只能用于函数内
	n, s := 0x1234, "hello ,world" //不同类型初始化值
	fmt.Println(&x, x, s, n)
	x, y := 233, "ccd" //此时对x的操作，退化为赋值操作，
	//x := 445 //至少有一个新变量被定义，no new variables on left side of :=
	fmt.Println(&x, x, y)
	//以上退化操作，对于重复使用err很有用
	//f, err := os.open("/dev/random")
	//...
	//n, err := f.read(buf)
	//多变量赋值时，先计算所有相关值，然后从左到右依次赋值
	data, i := [3]int{0, 1, 2}, 0
	i, data[i] = 2, 100 //(i = 0) -> (i = 2), (data[0] = 100)
	fmt.Println(data)   //[100 1 2]
	const z = 3423      // 代码块中未使用的常量不会报错
	fmt.Println("======================enumandConst=========================")
	enumandConst()
	fmt.Println("======================mathLearn=========================")
	mathLearn()
	fmt.Println("======================strconvLearn=========================")
	strconvLearn()
	fmt.Println("======================aliasLearn=========================")
	aliasLearn()
	fmt.Println("======================stringLearn=========================")
	stringLearn()
	fmt.Println("======================pointerLearn=========================")
	pointerLearn()
}

//iota是golang的常量计算器，只能在常量表达式中使用
func enumandConst() {
	const (
		x = iota //0
		y        //1
		z        //2
	)
	//每次遇到新的const关键字，iota计数都会重新开始
	const a = iota //0
	const (
		b = iota //0
		c        //1
	)

	const (
		i, j = 3, 1.1
		k    = iota //按行序自增 1
		k1          //2
		k2          //3
	)
	fmt.Println(k, k1, k2)
	const (
		_  = iota             // ignore first value by assigning to blank identifier
		KB = 1 << (10 * iota) // 1 << (10*1)
		MB                    // 1 << (10*2)
		GB                    // 1 << (10*3)
		TB                    // 1 << (10*4)
	)
	//自增作用范围为常量组。可在多常量定义中使用多个iota，它们各自单独计数
	//须确保组中每行常量的列数量相同
	const (
		_, _    = iota, iota * 10 //0, 0*10
		a1, a10                   //1 , 1 * 10
		b1, b10                   //2 , 2 * 10
	)
	const (
		i3 = iota //0
		j3 = 3.14 //1.23
		k3        //1.23
		l3 = iota //3
		m3        //4
		n3        //5
	)
	const xx int = 42
	//常量无法取地址
	//fmt.Println(&xx,xx) //cannot take the address	of xx
	const x3 = 100     //不做类型声明
	const y3 byte = x3 //直接将x3展开，相当于 const y3 byte = 100

	//const x4 int = 100	//显示进行类型声明，此时编译器会做强制类型检查
	//const y4 byte = x4 //此时编译会报错，cannot use x4 (type int) as type in const initializer
}

func mathLearn() {
	a, b, c := 11, 011, 0x11 //十进制， 八进制， 十六进制
	fmt.Println(a, b, c)
	fmt.Printf("%b, %#o, %#x", a, b, c) //分别以二进制， 八进制，十六进制显示，加#显示的更友好
	fmt.Println()
	fmt.Println(math.MinInt64, math.MaxInt64)
}

func strconvLearn() {
	a, _ := strconv.ParseInt("11001111", 2, 32)
	b, _ := strconv.ParseInt("0112", 8, 32)
	c, _ := strconv.ParseInt("11", 16, 32)
	d, _ := strconv.ParseInt("0x11", 0, 32)
	fmt.Println(a, b, c, d)
	fmt.Println("0b" + strconv.FormatInt(a, 2))
	fmt.Println("0" + strconv.FormatInt(a, 8))
	fmt.Println("0x" + strconv.FormatInt(a, 16))
}

type myint int

//即使是别名，在赋值时也需要显示类型转换
func aliasLearn() {
	//var a myint = 123
	//var b int = a //cannot use a (type myint) as type int in assignment
	//fmt.Println(a, b)
}

func stringLearn() {
	var ch byte = 65      //byte is a alias of int8
	var ch2 byte = '\x30' //十六进制表示
	fmt.Println(ch, ch2)
	fmt.Printf("ch:%c, ch2:%c\n", ch, ch2)
	var ch3 rune = '\u0041'                //小写u 2个字节
	var ch4 rune = '\U0002B7E7'            //大写U 4个字节
	fmt.Printf("%c,%c\n", ch3, ch4)        //以字符打印
	fmt.Printf("%U,%U,%U\n", ch, ch3, ch4) //以unicode码点打印
	rawString := `This is a raw string \n不起作用`
	fmt.Printf(rawString)
	var str string
	str += "str is not nil"
	fmt.Println(str)
	str = "你好"
	fmt.Printf("the length of %s is:%d\n", str, len(str)) //len求的是字节数
	str = "this is 一个测试 for strings"
	fmt.Println(strings.HasPrefix(str, "th"))
	fmt.Println(strings.HasSuffix(str, "gs"))
	fmt.Println(strings.Contains(str, "is"))
	fmt.Println(strings.Index(str, "个"))     //11
	fmt.Println(strings.LastIndex(str, "s")) //31
	fmt.Println(strings.IndexRune(str, '个')) //11
	fmt.Println(len("个"))                    //3
	fmt.Println(strings.Count(str, "s"))
	fmt.Println(strings.Repeat("LOVE", 3))
	fmt.Println(strings.ToLower(str)) //this is 一个测试 for strings
	fmt.Println(strings.ToUpper(str)) //THIS IS 一个测试 FOR STRINGS
	str_slice := strings.Split(str, " ")
	fmt.Println(str_slice)
	fmt.Println(strings.Join(str_slice, "-"))

}

func pointerLearn() {
	var i = 5
	fmt.Printf("the integer %d's address is:%p\n", i, &i)
	var intP *int
	intP = &i
	fmt.Printf("the value at memory address %p is %d\n", intP, *intP)
	s := "ciao"
	var p = &s
	*p = "good bye"
	fmt.Printf("now s is %s", s)

}
