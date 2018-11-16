package main

import "fmt"

func updateArray(s []int){
	s[0] = 100
}

func sliceOperation(){
	fmt.Println("test slice operation")
	arr1 := [...]int{0,1,2,3,4,5,6,7,8,9}
	slice1 := arr1[:]

	slice2 := slice1[:3]

	s3:= append(slice2, 11)
	fmt.Println(s3)
	fmt.Println(slice1)
	fmt.Println(arr1)

	//append longer than array
	fmt.Println("test append longer than array")
	s4 := append(slice1, 12)
	fmt.Println(s4)
	fmt.Println(slice1)
	fmt.Println(arr1)

}

func main() {
	arr1 := [...]int{0,1,2,3,4,5,6,7,8,9}

	a2 := arr1[2:6]
	fmt.Println(a2)
	fmt.Println(arr1[:6])
	fmt.Println(arr1[2:])
	updateArray(arr1[2:])
	fmt.Println(arr1[2:])
	fmt.Println(arr1[:])
	fmt.Println()

	sliceOperation()
	//切片的操作

}
