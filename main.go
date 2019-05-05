package main

import (
	"io"
	"os"
)

func main() {
	input, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	output, err := os.OpenFile(os.Args[1]+".mp3", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	b := make([]byte, 4096)
	for {
		n, err := input.Read(b)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		for i := 0; i < n; i++ {
			b[i] = b[i] ^ 0xa3
		}
		_, err = output.Write(b[0:n])
		if err != nil {
			panic(err)
		}
	}
	err = input.Close()
	if err != nil {
		panic(err)
	}
	err = output.Close()
	if err != nil {
		panic(err)
	}
}
