package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//byteData, err := os.ReadFile("go_first4/hello.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(byteData)) //可以看见文件内容
	//file, err := os.Open("go_first4/hello.txt") //打开文件
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//for {
	//	var byteData = make([]byte, 1204)
	//	n, err := file.Read(byteData) //读多少字节,n表示实际读了多少
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(string(byteData[:n]))
	//}
	file, err := os.Open("go_first4/hello.txt") //打开文件
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buf := bufio.NewScanner(file)
	buf.Split(bufio.ScanWords)
	var index int
	for buf.Scan() {
		index++
		fmt.Println(index, buf.Text())
	}

}
