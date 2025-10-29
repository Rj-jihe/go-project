package main

import (
	"fmt"
	"net"
	"time"
)

// 服务端
func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:81")
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("tcp server bind addr: ", addr.String())
	for {
		conn, err := listen.Accept()
		if err != nil {
			break
		}
		fmt.Println(conn.RemoteAddr())
		conn.Write([]byte("hello world"))
		time.Sleep(1 * time.Second)
		conn.Close()
	}
}
