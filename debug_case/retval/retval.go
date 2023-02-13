package main

import "fmt"

func foo(v int) *int {

	var f_v = v
	return &f_v
}

func main() {

	var m_v *int

	m_v = foo(100)

	fmt.Println(*m_v)

}
