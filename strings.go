package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	string1 := "Yes，我爱慕课网，!"
	fmt.Println("strng1 的字节长度",len(string1))

	fmt.Printf("%x\n",string1)

	for _,b := range []byte(string1){
		fmt.Printf("%x ",b)
	}
	fmt.Println("\n打印string range后的 %s 形式")
	for i,ch := range(string1){
		fmt.Printf("%d,%s ",i,ch)
	}

	fmt.Println("\n打印string1的二进制形式")

	for i,ch := range(string1){
		fmt.Printf("%d,%x ",i,ch)
	}
	fmt.Println("\n Rune Count string1:",utf8.RuneCountInString(string1))

	bytes := []byte(string1)

	//%c 打印 字符
	for len(bytes) > 0{
		ch,size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ",ch)
	}

	fmt.Println()
	for i,ch := range []rune(string1){
		fmt.Printf("%d,%c ",i,ch)
	}
	fmt.Println()
	test := strings.Fields(string1)
	fmt.Println(len(test),test)
	test = strings.Split(string1,"，")
	fmt.Println(len(test),test)
}
