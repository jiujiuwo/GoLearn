package main

import (
	"fmt"
	"regexp"
)

/*
	同样是内建的正则库，go语言的结果与java不同，不知道为什么。
	目前该方法是错误的，下一步寻找问题
 */
func isMatch(s string, p string) bool {
	isMatch,_ := regexp.MatchString(p,s)
	//isMatch,_ := regexp.Match(p,[]byte(s))	两个方法都是错的
	return isMatch
}

func main(){
	fmt.Println(isMatch("aa","a"))
}
