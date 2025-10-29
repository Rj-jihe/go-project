package main

import "fmt"

func main() {
	var userMap map[int]string = map[int]string{
		1: "枫枫",
		2: "张三",
		4: "",
	}
	//取值
	fmt.Println(userMap)
	fmt.Println(userMap[1])
	fmt.Println("%#v\n", userMap[4])
	value := userMap[4]
	value, ok := userMap[4]
	fmt.Println(value, ok)
	//设置值
	userMap[1] = "枫枫知道"
	fmt.Println(userMap)
	delete(userMap, 4)
	fmt.Println(userMap)
}
