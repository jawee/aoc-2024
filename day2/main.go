package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	// file, err := os.Open(pwd + "/day2/atest.txt")
	// file, err := os.Open(pwd + "/day2/btest.txt")
	file, err := os.Open(pwd + "/day2/input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	res := a(file)
	// res := b(file)
	fmt.Printf("%d\n", res)
}

func a(file io.Reader) int {
	panic("not implemented")
}

func b(file io.Reader) int {
	panic("not implemented")
}

