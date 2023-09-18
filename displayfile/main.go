package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	if len(args) < 1 {
		fmt.Println("file name missing")
		return
	}
	file, _ := os.Open("quest.txt")
	bb := make([]byte, 14)
	file.Read(bb)
	fmt.Println(string(bb))
}
