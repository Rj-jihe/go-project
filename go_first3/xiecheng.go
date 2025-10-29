package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup
var moneyChan1 = make(chan int) //声明并初始化一个长度为0的信道
// 变量钱信道，类型chan int
var nameChan1 = make(chan string)
var doneChan1 = make(chan struct{}) //声明一个用于关闭的信道

func send(name string, money int) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(5 * time.Second) //延时一秒钟
	fmt.Printf("%s 购物结束\n", name)
	moneyChan1 <- money
	nameChan1 <- name
	wait.Done() //表结束
}

// 协程(主要用于耗时函数）
func main() {
	startTime := time.Now() //开始购物时间
	//现在的模式有点像购物接力
	//shopping("rj")
	//shopping("ry")
	//主线程结束，协程函数跟着结束
	wait.Add(3) //有两个协程函数
	go send("rj", 2)
	go send("ry", 3)
	go send("rr", 5)
	go func() {
		defer close(moneyChan1)
		defer close(nameChan1)
		defer close(doneChan1)
		wait.Wait()

	}()
	var moneyList []int
	var nameList []string
	var event = func() {
		for {
			select {
			case money := <-moneyChan1:
				moneyList = append(moneyList, money)
			case name := <-nameChan1:
				nameList = append(nameList, name)
			case <-doneChan1:

				return
			}
		}
	}
}
event()
fmt.PrintIn("")