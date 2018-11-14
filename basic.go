package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

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


/*
	测试Go语言内建 复数类型
	以及复数库cmplx的使用
 */

 func variableComplex(){
 	fmt.Println("复数类型测试")
 	c := 3+4i
 	fmt.Println(c)
 	fmt.Println(cmplx.Abs(c))
 	fmt.Println( cmplx.Pow(math.E,1i*math.Pi) + 1)
 	fmt.Println(cmplx.Exp(1i*math.Pi)+1)
 }

 /*
 	强制类型转换演示
  */

func triangle(){
	fmt.Println("强制类型转换测试")
	var a,b int = 3,4
	var c int
	c = int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
}

func consts(){
	fmt.Println("常量测试")
	//普通值的枚举类型
	const filename = "abc.txt"
	const a,b = 3,4
	var c int
	c = int(math.Sqrt(a*a+b*b))
	fmt.Println(filename,c)
}
/*
	枚举类型（常量）
 */
func enums(){
	fmt.Println("枚举测试")
	//普通枚举类型
	const(
		cpp = 0
		java = 1
		python = 2
	)
	fmt.Println(cpp,java,python)

	//自增值枚举类型
	const(
		spring = iota
		summer
		fall
		winter
	)

	fmt.Println(spring,summer,fall,winter)

	const(
		b = 1 << (10*iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b,kb,mb,gb,tb,pb)

}

func main() {
	fmt.Println("Hello,Wolrd!")
	variablePackage()
	variableInitAuto()
	variableInitMan()
	variableTypeDeduction()
	variableTypeShort()
	variableComplex()
	triangle()
	consts()
	enums()
}
