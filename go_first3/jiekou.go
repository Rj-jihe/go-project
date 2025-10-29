package main

import "fmt"

type SingInterface interface {
	Sing()
}
type Chicken struct {
	Name string
}

func (c Chicken) Sing() {
	fmt.Println(c.Name, "在唱歌")
}

type Cat struct {
	Name string
}

func (c Cat) Sing() {
	fmt.Println(c.Name, "在唱歌")
}

// 唱歌
func sing(c SingInterface) {
	//ch, ok := c.(Chicken)
	//fmt.Println(ch, ok)//类型断言
	//枚举
	switch server := c.(type) {
	case Chicken:
		fmt.Println(server)
	case Cat:
		fmt.Println(server)
	default:
		fmt.Println("其他")
	}
	c.Sing()

}

func Print(val any) {
	fmt.Println(val)

}
func printAnything(v interface{}) {
	fmt.Printf("值: %v, 类型: %T\n", v, v)
}

func main() {
	a := Chicken{"jame"}
	b := Cat{"aimi"}
	sing(a)
	sing(b)
	Print(a)
	Print(b)
	printAnything(42)              // int
	printAnything("hello")         // string
	printAnything(Chicken{"jame"}) // main.Chicken
}
