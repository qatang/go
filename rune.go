package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

func main() {
	// 字符
	s := "阿达东澳岛"
	fmt.Println(len(s)) // 15

	fmt.Println()

	// utf-8编码，每个中午字符三个字节
	for _, b := range []byte(s) {
		fmt.Printf("%x ", b)
	}
	fmt.Println()

	for i, ch := range s {// ch is rune
		fmt.Printf("(%d %x)", i, ch) // 输出unicode
	}
	fmt.Println()

	fmt.Println("rune count : ", utf8.RuneCountInString(s))

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()

	fmt.Println(strings.Contains("aaa", "nnn"))
}
