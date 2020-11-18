package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	string1 := "Yes，我爱慕课网，!"
	//通过len函数获得的string的长度是字节的长度
	fmt.Println("strng1 的字节长度",len(string1))

	//按照字节的格式打印string
	fmt.Printf("%x\n",string1)

	for _,b := range []byte(string1){
		fmt.Printf("%x ",b)
	}
	//
	fmt.Println("\n打印string range后的 %s 形式")
	for i,ch := range(string1){
		fmt.Printf("%d,%s ",i,ch)
	}

	fmt.Println("\n打印string1的二进制形式")

	for i,ch := range(string1){
		fmt.Printf("%d,%x ",i,ch)
	}

	//正确的获取字符串的长度，rune为go语言中的一个字符
	fmt.Println("\n Rune Count string1:",utf8.RuneCountInString(string1))

	bytes := []byte(string1)

	//%c 打印 字符
	for len(bytes) > 0{
		ch,size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ",ch)
	}

	fmt.Println()

	//正确便利字符串的每个字符
	for i,ch := range []rune(string1){
		fmt.Printf("%d,%c ",i,ch)
	}


	fmt.Println()
	//将字符串按照空格进行切割
	test := strings.Fields(string1)
	fmt.Println(len(test),test)
	//将字符串按照任意字符进行切割
	test = strings.Split(string1,"，")
	fmt.Println(len(test),test)
}
