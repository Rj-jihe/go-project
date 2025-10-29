package main

import "fmt"

func main() {
	fmt.Printf("请输入你的名字：")
	var name string
	fmt.Scan(&name)
	fmt.Println(name)
	fmt.Printf("请输入你的年龄：")
	var age int
	n, err := fmt.Scan(&age)
	fmt.Println(n, err, age)
}
