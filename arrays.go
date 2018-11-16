package main

import "fmt"

func testArray1(){

	//数组定义
	var arr1 [5]int
	arr2 := [3]int{1,2,3}
	arr3 := [...]int{2,4,6,8}
	fmt.Println(arr1,arr2,arr3)

	var grid [4][5]int
	fmt.Println(grid)

	//遍历数组
	for i:= 0;i < len(arr2);i++{
		fmt.Println(arr2[i])
	}

	for i,v := range arr3{
		fmt.Println(i,v)
	}
}

func printArray(arr [5]int){
	fmt.Println("print array:")
	for i,v := range arr{
		fmt.Println(i,v)
	}
	arr[0] = 100
}

func printArrayPtr(arr *[5]int){
	fmt.Println("print array:")
	for i,v := range arr{
		fmt.Println(i,v)
	}
	arr[0] = 100
}

func main() {
	testArray1()
	var arr1 [5]int
	//arr2 := [3]int{1,2,3}
	//arr3 := [...]int{2,4,6,8}
	printArray(arr1)
	//printArray(arr2)出错
	printArray(arr1)
	printArrayPtr(&arr1)
	printArray(arr1)
}
