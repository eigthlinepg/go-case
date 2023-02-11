package main

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	tt := GetInstance()
	fmt.Println(tt)

}
