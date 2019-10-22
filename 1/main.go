package main

import (
	"bytes"
	"fmt"
	"os"
)

var encoded_flag = []byte{23, 29, 1, 92, 3, 7, 25, 0, 30, 7, 1, 17, 12, 3, 52, 28, 42, 26, 3, 13, 46, 10, 27, 11, 2, 46, 7, 31, 11, 25}
var key = []byte{116, 104, 101, 107, 101, 121}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: <flag>")
		os.Exit(1)
	}

	flag := make([]byte, 0)
	for i := 0; i < len(os.Args[1]); i++ {
		flag = append(flag, os.Args[1][i]^byte(i)^key[i%len(key)])
	}

	if bytes.Equal(flag, encoded_flag) {
		fmt.Println("Yes your flag!")
	} else {
		fmt.Println("No no nope, just isn't the flag")
	}
}
