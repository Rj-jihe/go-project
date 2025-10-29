package main

import (
	"fmt"
	"reflect"
)

func setValue(obj any, value any) {
	v1 := reflect.ValueOf(obj)

	// 检查obj是否是指针类型
	if v1.Kind() != reflect.Ptr {
		fmt.Println("错误：第一个参数必须是指针")
		return
	}

	// 检查指针指向的值是否可设置
	if !v1.Elem().CanSet() {
		fmt.Println("错误：值不可设置")
		return
	}

	v2 := reflect.ValueOf(value)

	// 检查类型是否匹配
	if v1.Elem().Kind() != v2.Kind() {
		fmt.Printf("错误：类型不匹配，期望 %v，实际 %v\n", v1.Elem().Kind(), v2.Kind())
		return
	}

	// 根据类型设置值
	switch v1.Elem().Kind() {
	case reflect.String:
		v1.Elem().SetString(v2.String())
		fmt.Println("成功设置字符串值")
	case reflect.Int:
		v1.Elem().SetInt(v2.Int())
		fmt.Println("成功设置整数值")
	default:
		fmt.Printf("不支持的类型: %v\n", v1.Elem().Kind())
	}
}

func main() {
	var name = "枫枫"
	var age = 24

	fmt.Println("修改前:", name, age)

	setValue(&name, "张十三")
	setValue(&age, 18)

	fmt.Println("修改后:", name, age)

	// 测试错误情况
	setValue(name, "错误示例") // 非指针参数
	setValue(&age, "字符串")  // 类型不匹配
}
