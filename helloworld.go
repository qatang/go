package main

import "fmt"

func main() {
	fmt.Println("hello world")

	a, b := 3, 4
	fmt.Printf("内存地址：pa=%x\n值：b=%x", &a, b)
}
