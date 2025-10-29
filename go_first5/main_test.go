package main

import (
	"fmt"
	"os"
	"testing"
)

// 测试前
func setup() {
	fmt.Println("测试前")
}

// 测试后
func teardown() {
	fmt.Println("测试后")
}

func TestAdd2(t *testing.T) {
	fmt.Println("测试中")
}

// 测试的主函数
func TestMain(m *testing.M) {
	fmt.Println("testMain")
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
