package main

import (
	"fmt"
	"reflect"
)

func getType(obj any) {
	t := reflect.TypeOf(obj)
	switch t.Kind() {
	case reflect.String:
		fmt.Println("string")
	case reflect.Int:
		fmt.Println("int")
	case reflect.Struct:
		fmt.Println("struct")
	}
}
func main() {
	getType("12")
	getType(123)
	getType(struct {
		Name string
	}{})
}
