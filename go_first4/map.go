package main

import (
	"fmt"
	"sync"
)

func main() {
	var maps = sync.Map{}
	go func() {
		for {
			maps.Store("a", "张三")
		}
	}()
	go func() {
		for {
			val, ok := maps.Load("a")
			fmt.Println(val, ok)
		}
	}()
	select {}
}
