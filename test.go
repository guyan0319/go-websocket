package main

import (
	"fmt"
)

type User struct {
	Id     int
	Name   string
	Amount float64
}


func main() {
	var i interface{}
	i = "string"
	fmt.Println(i)
	i = 1
	fmt.Println(i)
	i = User{Id: 2}
	//i.(User).Id = 15  //运行此处会报错，在函数中修改interface表示的结构体的成员变量的值，编译时遇到这个编译错误，cannot assign to i.(User).Id
	fmt.Println(i.(User).Id)
	i = &User{Id: 2}
	i.(*User).Id = 15
	fmt.Println(i.(*User).Id)
}