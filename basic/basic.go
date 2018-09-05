package basic

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func variableTypeDeduction() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	// 类型转换是强制的
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

const filename = "abc.txt"

func consts() {
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)

	const (
		cpp1    = iota
		_
		java1   = 1
		python1 = 2
		golang1 = 3
	)

	// b, kb, mb, gb, tb, pb
	const(
		b = 1 << (10 * iota)
		// 下面依次也会执行
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableInitialValue()
	variableShorter()
	fmt.Println(aa, ss, bb)
	euler()
	triangle()
	consts()
	enums()
}
