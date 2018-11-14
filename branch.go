package main

import (
	"fmt"
	"io/ioutil"
)

const filename = "abc.txt"

/*
	基础if语句
 */
func testIf1(){
	fmt.Println("基础if语句测试")
	contents, err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n",contents)
	}
}

/*
	if后可以赋值，作用域为if块
 */
func testIf2(){
	fmt.Println("带赋值的if语句测试")
	if contents,err := ioutil.ReadFile(filename); err ==nil{
		fmt.Printf("%s\n",contents)
	}else {
		fmt.Println(err)
	}
}

/*
	switch语句新写法
 */

 func testSwicthScore(score int) string{
 	fmt.Println("新的switch写法")
 	g := ""
	 switch {
	 case score <0||score >100:
	 	panic(fmt.Sprintf("Wrong score: %d",score))
	 case score < 60:
	 	g = "F"
	 case score < 70:
		g = "C"
	 case score < 90:
	 	g = "B"
	 case score <= 100:
	 	g = "A"
	 }
 	
 	return g
 }

 func testSwitchBasic(a,b int,op string) int{

 	fmt.Println("基础switch语句测试")

 	var result int

	 switch op {
	 case "+":
	 	result = a + b
	 case "-":
	 	result = a - b
	 case "*":
	 	result = a * b
	 case "/":
	 	result = a /b
	 default:
		 panic("unsupport operator: "+op)

	 }

 	return result
 }

func main() {
	testIf1()
	testIf2()
	fmt.Println(testSwitchBasic(2,3,"+"))
	fmt.Println(testSwicthScore(0),testSwicthScore(59),testSwicthScore(100))
}
