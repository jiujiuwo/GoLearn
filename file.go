package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func createFile() {
	file1, err := os.Create("file1.txt")
	if err != nil {
		panic(err)
	}
	file1.WriteString("file1 write string test!")
	file1.Close()
}

func readFile1() {
	file1, err := ioutil.ReadFile("file1.txt")
	if err != nil {
		panic(err)
	}
	tmp := string(file1)
	fmt.Println(tmp)
}

func readFile2() {
	file1, err := os.Open("file1.txt")
	if err != nil {
		panic(err)
	}
	var result [1024]byte
	_, err = file1.Read(result[:])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result[:]))
}

func main() {
	createFile()
	readFile1()
	readFile2()
}
