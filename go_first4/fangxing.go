package main

import (
	"encoding/json"
	"fmt"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func plus[T int | uint](n1, n2 T) T {
	return n1 + n2
}
func myPrint[T int, K string](n1 T, n2 K) {

}

//	func plus[T Number](n1, n2 T) T {
//		return n1 + n2
//	}
type Response struct { //jason的反解
	Code int    `json:"code"`
	Myg  string `json:"myg"`
	Data any    `json:"data"`
}

func main() {
	//plus(1, 2)
	//var u1, u2 = uint(2), uint(3)
	//plus(u1, u2)
	type User struct {
		Name string `json:"name"`
	}
	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	user := Response{
		Code: 0,
		Myg:  "成功",
		Data: User{
			Name: "枫枫",
		},
	}
	byteData, _ := json.Marshal(user)
	fmt.Println(string(byteData))
	userInfo := Response{
		Code: 0,
		Myg:  "成功",
		Data: UserInfo{
			Name: "枫枫",
			Age:  20,
		},
	}
	byteData, _ = json.Marshal(userInfo)
	fmt.Println(string(byteData))
}
