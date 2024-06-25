package main

import (
	."fmt"
	"reflect"
)

func main() {
	num := 1
	var num1 int32 = 123123

	Println(num)
	Println(num1)
	Println(reflect.TypeOf(num))
	Println(reflect.TypeOf(num1))
}