package main

import "fmt"

func main() {
	defer fmt.Println("defer2")
	defer fmt.Println("defer1")
	return
}
