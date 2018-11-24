package main

import (
	"fmt"
	"strconv"
)
//哇，go语言同样的解法，beat 100%
func isPalindrome(x int) bool {
	tmpX := strconv.Itoa(x)
	len := len(tmpX)
	for index,_ := range []byte(tmpX){
		if(tmpX[index]==tmpX[len - index -1]){
			continue
		}else{
			return false
		}
	}
	return true
}

func main(){
	fmt.Println(isPalindrome(818))
}