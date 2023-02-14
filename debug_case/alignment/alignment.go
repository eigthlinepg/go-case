package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Id       int64  //8
	Name     string //16
	Age      int8   //1
	IsActive bool   //1
}

func main() {
	var u User
	fmt.Println("Size of user struct:", unsafe.Sizeof(u))

	fmt.Println("Offset of ID:", unsafe.Offsetof(u.Id))
	fmt.Println("Offset of Name:", unsafe.Offsetof(u.Name))
	fmt.Println("Offset of Age:", unsafe.Offsetof(u.Age))
	fmt.Println("Offset of IsActive:", unsafe.Offsetof(u.IsActive))
}
