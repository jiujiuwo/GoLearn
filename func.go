package main

import (
	"fmt"
	"reflect"
)

func testFunc(a,b int,op string) int{

	switch op{
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		//return a / b
	default:
		panic("unsupported operation: "+op)
	}
	return 0
}
/*
	多返回值一般用于返回错误信息
 */
func testFuncWithErr(a,b int,op string) (int,error){

	switch op{
	case "+":
		return a + b,nil
	case "-":
		return a - b,nil
	case "*":
		return a * b,nil
	case "/":
		//return a / b
		q,_ := funcTestDiv(a,b)
		return q,nil
	default:
		return 0,fmt.Errorf("unsupported operation: %s",op)
	}
}

/*
	函数多返回值
 */
func funcTestDiv(a,b int)(int, int ){
	return a/b,a%b
}

/*
	多返回值及名字

*/
func funcTestDivName(a,b int)(c,d int){
	return a/b,a%b
 }

/*
	函数式编程
 */

 func apply(op func(int,int) int, a,b int) int{
 	reflect.ValueOf(op).Pointer()
 	fmt.Printf("Calling function %s witg args (%d,%d)")
 	return op(a,b)
 }

func main() {
	fmt.Println(testFunc(1,2,"+"))
	fmt.Println(funcTestDiv(5,7))
	c,d := funcTestDivName(6,7)
	fmt.Println(c,d)

	fmt.Println(testFuncWithErr(1,2,"x"))

	if result,err := testFuncWithErr(3,4,"a");err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}
}
