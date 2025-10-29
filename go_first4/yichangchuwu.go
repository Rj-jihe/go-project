package main

import (
	"fmt"
	"runtime/debug"
)

// import (
//
//	"errors"
//	"fmt"
//
// )
//
//	func div(a, b int) (res int, err error) {
//		if b == 0 {
//			err = errors.New("除数不能为0") //创建一个错误将其赋给err
//			return
//		}
//		res = a / b
//		return
//	}
//
//	func server() (res int, err error) {
//		res, err = div(2, 0)
//		if err != nil {
//			//把错误向上传递
//			return
//		}
//		//执行其他的逻辑
//		res++
//		res += 2
//		return
//	}
//
//	func main() {
//		res, err := server()
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println(res)
//
// }
// 中断程序
//
//	func init() {
//		_, err := os.ReadFile("123456")
//		if err != nil {
//			panic("错误了")
//			//log.Fatalln("错误")
//		}
//	}
//
// func main() {
//
//		fmt.Println("main")
//	}
//
// 异常捕获
func read() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			//打印堆栈
			fmt.Println(string(debug.Stack())) //找错误的地方
		}
	}()
	var list []int = []int{1, 2}
	fmt.Println(list[2])
}
func main() {
	read()
	//正常逻辑
	fmt.Println("正常逻辑")
}
