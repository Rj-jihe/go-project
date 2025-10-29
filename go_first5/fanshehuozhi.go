package main

import (
	"fmt"
	"reflect"
)

func getValue(obj any) {
	v := reflect.ValueOf(obj)
	switch v.Kind() {
	case reflect.String:
		fmt.Println("string", v.String())
	case reflect.Int:
		fmt.Println("int", v.Int())
	case reflect.Struct:
		fmt.Println("struct", v.Type())
	}
}
func main() {
	getValue("12")
	getValue(123)
	getValue(struct {
		Name string
	}{})
}
