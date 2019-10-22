package main

import (
	"bytes"
	"fmt"
	"os"
)

var encoded_flag = []byte{__ENCODED__}
var key = []byte{__KEY__}

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
