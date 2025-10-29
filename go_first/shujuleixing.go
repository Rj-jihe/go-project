package main

import "fmt"

func main() {
	//var age = 12
	//var u8 uint8 = 255
	//nint8
	//0 0 0 0 0 0 0 0=0
	//1 1 1 1 1 1 1 1=255=2^8-1
	//int8
	//0    1 1 1 1 1 1 1=2^7-1
	//1    0 0 0 0 0 0 1=-1
	//1    1 1 1 1 1 1 1 =-128
	var a byte = 'a' //ascii里面的字符
	fmt.Printf("%c %d\n", a, a)
	var a1 uint8 = 97
	fmt.Printf("%c %d\n", a1, a1)
	//
	var z rune = '中'
	fmt.Printf("%c %d\n", z, z)
	fmt.Println("枫枫\t知道")
	fmt.Println("'枫枫'知道")
	fmt.Println("\"枫枫\"知道")
	fmt.Println("\\枫枫\\知道")
	fmt.Println("枫枫\n知道")
	fmt.Println("枫枫\r知道")
	fmt.Println(`hello \n
world`)
	var a5 int
	var a2 float32
	var a3 bool
	var a4 string
	fmt.Printf("%#v\n", a5)
	fmt.Printf("%#v\n", a2)
	fmt.Printf("%#v\n", a3)
	fmt.Printf("%#v\n", a4)
}
