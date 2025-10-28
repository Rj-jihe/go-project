package main

import (
	"fmt"
	"study/go_first/version"
)

func hello() {
	fmt.Println(age)
	fmt.Println("Hello World")
}

var age = 12

const Version string = "2.0.1"

func main() {
	fmt.Println(version.Version)
	//先声明
	var name string
	name = "Jack"
	fmt.Println(name)

	//声明并赋值
	var name1 = "Jack"
	fmt.Println(name1)
	//省略类型
	var name2 = "Jack"
	fmt.Println(name2)

	name3 := "Jack"
	fmt.Println(name3)

	hello()

	var a1, a2 = 1, 2
	fmt.Println(a1, a2)

}
