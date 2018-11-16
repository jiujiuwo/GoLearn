package main

import "fmt"

func updateArray(s []int){
	s[0] = 100
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


}
