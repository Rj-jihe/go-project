package main

import (
	"fmt"
	"os"
)

func main() {
	//flag 文件打开方式
	//file, err := os.OpenFile("go_first4/W.txt", os.O_CREATE|os.O_RDWR, 0770)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	////file.Write([]byte("Hello World")) //写文件内容
	//byteData, err := io.ReadAll(file)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(byteData))
	//err := os.WriteFile("go_first/w1.txt", []byte("hello world"), 0777)
	//fmt.Println(err)
	//rFile, err := os.Open("C:\\Users\\j5753\\Pictures\\微信图片_20251017093428.jpg")
	//if err != nil {
	//	panic(err)
	//}
	//defer rFile.Close()
	//wFile, err := os.OpenFile("go_first4/w2.jpg", os.O_CREATE|os.O_WRONLY, 0)
	//if err != nil {
	//	panic(err)
	//}
	//defer wFile.Close()
	//io.Copy(wFile, rFile)
	dir, err := os.ReadDir("go_first4")
	if err != nil {
		panic(err)
	}
	for _, entry := range dir {
		info, _ := entry.Info()
		fmt.Println(entry.IsDir(), entry.Name(), info.Size())
	}
}
