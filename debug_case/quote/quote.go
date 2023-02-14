package main

import (
	"fmt"
	"reflect"
)

type open interface {
	/* TODO: add methods */
}

func main() {

	num := 100
	ptr := &num
	ptr2 := &num

	fmt.Println("num", num)
	fmt.Println("&num", &num)
	fmt.Println("ptr", ptr)
	fmt.Println("*ptr", *ptr)

	if ptr == ptr2 {
		fmt.Println("ptr equal ptr2")
	}

	//map
	colors := map[string]string{
		"red":   "#f0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Printf("colors %T\n", colors)

	//chan
	ch := make(chan int, 1)
	fmt.Printf("ch %T\n", ch)

	s := []int{1, 2, 3}
	fmt.Printf("ch %T\n", s)

	//interface
	var x interface{}
	x = 42
	t := reflect.TypeOf(x)
	fmt.Println("x=42, type:", t)

	x = "hello"
	t = reflect.TypeOf(x)
	fmt.Println("x=hello type:", t)

}
