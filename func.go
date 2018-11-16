package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
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
		return a / b
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
	函数式编程，函数作为形参（复合函数）
 */

 func apply(op func(int,int) int, a,b int) int{

 	p := reflect.ValueOf(op).Pointer()
 	opName := runtime.FuncForPC(p).Name()

 	fmt.Printf("Calling function %s with args (%d,%d)\n",opName,a,b)
 	return op(a,b)
 }

 func myPow(a,b int) int {
	 return int(math.Pow(float64(a),float64(b)))
 }

 /*
 	可变参数列表
  */
func mySum(numbers ...int) int{
	s := 0
	//range 的用法
	for i := range numbers{
		s += numbers[i]
	}
	return s
}

func swap(a,b int) {
	a,b = b,a
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

	//反射和函数作为参数
	fmt.Println(apply(myPow,3,4))

	fmt.Println(apply(func(a,b int) int{
		return int(math.Pow(float64(a),float64(b)))
	},3,4))

	fmt.Println(mySum(1,3,5,7,9,11))
}
