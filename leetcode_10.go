package main

import (
	"fmt"
	"regexp"
)

/*
	同样是内建的正则库，go语言的结果与java不同，不知道为什么。
 */
func isMatch(s string, p string) bool {
	isMatch,_ := regexp.MatchString(p,s)
	return isMatch
}

func main(){
	fmt.Println(isMatch("aa","a"))
}
