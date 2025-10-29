package main

import "fmt"

// 1。自定义类型
type MyInt int

// 2。类型别名
type NemInt = int

type MyCode int
type MylistCode = int

func (m MyCode) Code() {

} ///自定义类型可以仿方法，类型别名不可以
const a MyCode = 1
const b MylistCode = 1

func main() {
	var i MyInt
	var n NemInt
	fmt.Printf("type:%T value:%v\n", i, i)
	fmt.Printf("type:%T value:%v\n", n, n)
	fmt.Printf("type:%T value:%v\n", a, a)
	fmt.Printf("type:%T value:%v\n", b, b)
}
