package main

import "fmt"

func main() {
	var age int
	fmt.Printf("请输入你的年龄：")
	fmt.Scan(&age)
	//中断式
	if age <= 0 {
		fmt.Println("未出生")
		return
	}
	if age <= 18 {
		fmt.Println("未成年")
		return
	}
	if age <= 35 {
		fmt.Println("青年")
		return
	}

	fmt.Println("中年")
	//嵌套
	//多条件（&& || ！）

}
