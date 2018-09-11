package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// 函数要点
// 1. 返回值类型写在最后面
// 2. 可返回多个值
// 3. 函数作为参数
// 4. 没有默认参数、可选参数等，只有可变参数列表

// 值传递和引用传递
// cpp中可以值传递也可以引用传递（&）
// python和java除了系统内建类型是值传递以外，都是引用传递
// golang只有值传递，要改变值必须用指针参数

// 下划线表示变量不想用，其他变量名定义了不用会编译不过
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

// 两个返回值的三种写法
// 1.
// 13 / 3 = 4 ... 1
//func div(a, b int) (int, int) {
//  return a / b, a % b
//}

// 2. 函数返回多个值时可以起名字，对于调用者而言没有区别
func div(a, b int) (q, r int) {
	return a / b, a % b
}

// 3. 也可以，但不推荐。仅用于非常简单的函数
//func div(a, b int) (q, r int) {
//  q = a / b
//  r = a % b
//  return
//}

// 函数作为参数
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args "+
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println("Error handling")
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Printf("13 div 3 is %d mod %d\n", q, r)

	fmt.Println("pow(3, 4) is:", apply(
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))

	fmt.Println("1+2+...+5 =", sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println("a, b after swap is:", a, b)
}
