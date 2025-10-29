package main

import "fmt"

func main() {
	var nameList [3]string = [3]string{"枫枫", "张三", "李四"}
	fmt.Println(nameList)
	//索引
	fmt.Println(nameList[2])
	fmt.Println(nameList[len(nameList)-1])
	nameList[0] = "RJ"
	fmt.Println(nameList)
}
