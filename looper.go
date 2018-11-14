package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
/*
	基础循环，测试for，十进制转二进制
 */
func convertToBin(n int) string{
	result := ""
	for ;n > 0; n /= 2{
		lsb := n % 2
		result = strconv.Itoa(lsb)+result
	}
	return result
}

/*
	测试for读文件
 */
func printFile(filename string){

	fmt.Println("测试 printFile")

	file,err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	//没有起始条件和递增条件，分号可以省略，相当于while,
	//go语言中没有while
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func forever(){
	for{
		fmt.Println("abc")
	}
}

func main() {

	fmt.Println(convertToBin(5),convertToBin(13),convertToBin(72387885))
	printFile("abc.txt")
	forever()
}
