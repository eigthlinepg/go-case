package main

import "fmt"

func main() {

	var ilen int

	i := 34

	fmt.Println("addr ilen:", &ilen)
	fmt.Println("addr i:", &i)

	p := new(int)
	*p = 42

	fmt.Println("addr p:", &p)

}
