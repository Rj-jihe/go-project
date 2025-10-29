package main

import (
	"fmt"
	"sort"
)

func main() {
	var name []string
	//name = append(name, "RJ")
	//name = append(name, "张三")
	//name = append(name, "李四")
	fmt.Println(name == nil)
	//name = append(name, "人")
	//fmt.Println(name)
	//fmt.Println(name[0])
	ageList := make([]int, 3)
	fmt.Println(ageList)

	array := [3]int{1, 2, 3}
	slices := array[:]
	fmt.Println(slices)
	fmt.Println(array[0:2])

	var ints = []int{4, 8, 3, 2}
	sort.Ints(ints)
	fmt.Println(ints)                            //升序
	sort.Sort(sort.Reverse(sort.IntSlice(ints))) //降序
	fmt.Println(ints)
}
