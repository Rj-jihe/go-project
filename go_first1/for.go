package main

import "fmt"

func main() {
	var sum int = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	//死循环
	//for i := 0; true; i++ {
	//	fmt.Println(time.Now())
	//	time.Sleep(1 * time.Second)
	//}
	//for true {
	//		fmt.Println(time.Now())
	//	time.Sleep(1 * time.Second)
	var sum1 int = 0
	var i int = 0
	for i <= 100 {
		sum1 += i
		i++
	}
	fmt.Println(sum1)
	var list = []string{"枫枫", "知道"}
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
	for index, item := range list {
		fmt.Println(index, item)
	}
	//var userMap= map[string]string{"name", "知道"}
	//for key, value := range userMap {
	//	fmt.Println(key, value)
	//}
	//乘法表
	//for i := 1; i <= 9; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%d * %d = %d\t", i, j, i*j)
	//	}
	//	fmt.Println()
	//
	//}
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)

	}

}
