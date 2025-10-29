package main

import "fmt"

type UserInfo struct {
	Name string `json: "name"`
}

func (u *UserInfo) SetName(name string) {
	u.Name = name
}
func main() {
	user := UserInfo{Name: "枫枫"}
	user.SetName("张三")
	fmt.Println(user)
}
