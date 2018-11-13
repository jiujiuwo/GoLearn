package main

import "fmt"

var(
	aa = 1
	bb = "test"
	cc = false
)

func variablePackage(){
	fmt.Println(aa,bb,cc)
}

func variableInitAuto(){
	var a int
	var s string
	fmt.Printf("%d,%q\n",a,s)
}

func variableInitMan(){
	var a,b int = 3,4
	var s string = "abc"
	fmt.Println(a,b,s)
}

func variableTypeDeduction(){
	var a,b,s= 3,4,"edg"
	fmt.Println(a,b,s)
}

func variableTypeShort(){
	a,b,s := 3,4,true
	fmt.Println(a,b,s)
}

func main() {
	fmt.Println("Hello,Wolrd!")
	variablePackage()
	variableInitAuto()
	variableInitMan()
	variableTypeDeduction()
	variableTypeShort()
}
