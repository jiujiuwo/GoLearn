package main

import (
	"fmt"
)

func testCreateSlice(){
	fmt.Println("testCreateSlice")
	slice1 := []int{2,4,6,8}
	fmt.Println(slice1)
	slice2 := make([]int,16)
	fmt.Println(slice2)
	slice3 := make([]int,10,32)
	fmt.Println(slice3)
}

func testCopySlice(){
	fmt.Println("test copy slice")
	slice1 := []int{1,2,3,4,5,6}
	s2 := make([]int,16)
	s3 := make([]int,2)
	copy(s2,slice1)
	fmt.Println(s2)
	fmt.Println(slice1)
	copy(s3,slice1)
	fmt.Println(s3)

}

func testDeleteSliceElements(){
	fmt.Println("test Delete Slice Elements")
	slice1 := []int{1,2,3,4,5,6}
	s2 := append(slice1[:2],slice1[3:]...)
	fmt.Println(s2)
}

func testSlicePop(){
	fmt.Println("test slice pop")
	slice1 := []int{1,2,3,4,5,6}
	//front := slice1[0]
	slice1 = slice1[1:]
	fmt.Println(slice1)
	slice1 = slice1[0:len(slice1)-1]
	fmt.Println(slice1)
}

func testSliceLenAndCap(){
	fmt.Println("testSliceLenAndCap")
	var s []int //定义一个类型为int的slice
	fmt.Println(s)
	for i:=0;i<100;i++{
		fmt.Println(len(s), cap(s))
		s = append(s, i)
	}
	fmt.Println(s)
}

func main() {

	testSliceLenAndCap()
	testCreateSlice()
	testCopySlice()
	testDeleteSliceElements()
	testSlicePop()
}
