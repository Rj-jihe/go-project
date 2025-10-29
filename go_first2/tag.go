package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string
	Age      int
	Password string
}

func main() {
	user := User{Name: "枫枫", Age: 18, Password: "123456"}
	byteData, _ := json.Marshal(user)
	fmt.Println(string(byteData))
}
