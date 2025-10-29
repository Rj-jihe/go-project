package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name1 string `big:"-"`
	Name2 string
}

func (User) Call(name string) {
	fmt.Println("我被调用了", name)
}
func Call(obj any) {
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < v.NumMethod(); i++ {
		m := t.Method(i)
		if m.Name != "Call" {
			continue
		}
		method := v.Method(i)
		method.Call([]reflect.Value{
			reflect.ValueOf("枫枫"),
		})
	}
}

//读取标签
//type Student struct {
//	Name  string `json:"name"`
//	Age   int
//	IsMan bool
//}

//func ParseJson(obj any) {
//	v := reflect.ValueOf(obj)
//	t := reflect.TypeOf(obj)
//	for i := 0; i < v.NumField(); i++ {
//		tf := t.Field(i) //拿到了属性
//		jsonTag := tf.Tag.Get("json")
//		if jsonTag == "" {
//			jsonTag = tf.Name
//		}
//
//		fmt.Println(jsonTag, v.Field(i))
//	}
//}

func main() {
	s := User{} //将name改为大写
	Call(&s)

}
