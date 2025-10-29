package main

import "fmt"

func sayhello() {
	fmt.Println("hello")
}
func param1(id string) {
	fmt.Println("id:", id)
}
func param2(id, user string) {
	fmt.Println("id:,", id)
}
func r1() {
	return //无返回值
}
func r2() bool {
	return true
}

// r3

func r3() (string, bool) {
	if 1 > 2 {
		return "", false
	}
	if 1 > 2 {
		return "", false
	}
	return "", true
}

func main() {
	sayhello()
	param1("123")
	//匿名函数
	//var getName = func() string {
	//	return "枫枫"
	//}
	//var setName = func(name string) {
	//	fmt.Println(name)
	//	return
	//}
	//fmt.Println(getName())
	//setName("张三")
	fmt.Println("请输入你要执行的操作:")
	fmt.Println("1 登录")
	fmt.Println("2 注册")
	fmt.Println("3 个人中心")
	var index int
	fmt.Scan(&index)
	var funMap = map[int]func(){
		1: login,
		2: reser,
		3: userCenter,
	}
	fun, ok := funMap[index]
	if ok {
		fun()
	}
	//switch index {
	//case 1:
	//	login()
	//case 2:
	//	reser()
	//case 3:
	//	userCenter()
	//}
}
func login() {
	fmt.Println("登录")

}
func reser() {
	fmt.Println("注册")
}
func userCenter() {
	fmt.Println("个人中心")
}
