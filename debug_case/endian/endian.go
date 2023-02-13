package main

import (
	"encoding/binary"
	"fmt"
)

func main() {

	var i int32 = 0x01020304

	b := make([]byte, 4)

	binary.LittleEndian.PutUint32(b, uint32(i))

	if b[0] == 0x40 {
		fmt.Println("LittleEndian")
	} else {
		fmt.Println("BigEndian")
	}

}
