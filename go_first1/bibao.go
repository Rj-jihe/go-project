package main

import (
	"fmt"
	"time"
)

func awaitAdd(awaitSecond int) func(...int) int {
	time.Sleep(time.Duration(awaitSecond) * time.Second)
	return func(numberList ...int) (sum int) {

		for _, i2 := range numberList {
			sum += i2

		}
		return sum
	}
}

func copy(name string) {
	fmt.Printf("infun %p\n", &name)
	name = "枫枫不知道"
}
func set(name *string) { //name *指内存地址 {
	fmt.Printf("%p\n", name)
	//通过内存地址去找值(解引用）
	*name = "枫枫不知道"
}
func main() {
	//t1 := time.Now()
	//sum := awaitAdd(2)(1, 2, 3)
	//subTime := time.Since(t1)
	//fmt.Println(sum, subTime)
	var name = "枫枫"
	fmt.Printf("%p\n", &name)
	copy(name)
	//fmt.Println(name)
	set(&name) //传内存地址
	fmt.Println(name)
}
