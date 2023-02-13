package main

import "fmt"

func main() {

	//go build -gcflags "-N -l" -v
	a := sum(1, 2)
	fmt.Println(a)
}

func sum(i, j int) int {
	return i + j
}
