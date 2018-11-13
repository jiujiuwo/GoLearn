package main

import "fmt"

// 变量统一简化定义
var(
	aa = 1
	bb = "test"
	cc = false
)

/*
	包内变量
 */
func variablePackage(){
	fmt.Println(aa,bb,cc)
}

/*
 	变量定义及默认值，go语言中所有变量都有默认，确定的值
 */
func variableInitAuto(){
	var a int
	var s string
	fmt.Printf("%d,%q\n",a,s)
}

/*
	基础变量定义及赋初值
 */
func variableInitMan(){
	var a,b int = 3,4
	var s string = "abc"
	fmt.Println(a,b,s)
}

/*
	多变量同时赋值
 */
func variableTypeDeduction(){
	var a,b,s= 3,4,"edg"
	fmt.Println(a,b,s)
}

/*
:= 变量定义及赋值语法
 */
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
