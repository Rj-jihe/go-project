package main

import "fmt"

func main() {
	fmt.Println("Hello", "你好")
	fmt.Println("你好", "枫枫")
	fmt.Printf("%s哇，你好美!\n", "枫枫")
	fmt.Printf("%s哇，你好美!\n", "枫枫")
	fmt.Printf("%d\n", 2)
	fmt.Printf("%.2f\n", 2.455)
	fmt.Printf("%T %T\n", "你好", 2.5)
	fmt.Printf("%v\n", 2.5)
	fmt.Printf("%#v\n", " ")
	var f = fmt.Sprintf("%.2f", 2.455)
	fmt.Println(f)
}
