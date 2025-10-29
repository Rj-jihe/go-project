package main

import "fmt"

type Animal struct {
	Name string
}

type Cat struct {
	Animal
}

func main() {
	var a = Animal{Name: "猫"}
	cat1 := Cat{
		Animal: a,
	}
	fmt.Println(cat1.Name)
}
