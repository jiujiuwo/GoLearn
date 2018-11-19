package main

import (
	"fmt"
	"time"
)

func longestPalindrome(s string) string {
	result := ""
	tmp := ""
	for i:= 0;i < len(s);i++{
		for j:=len(s);j >i;j--{
			tmp = s[i:j]
			if isPalindrome(tmp){
				if len(tmp) > len(result){
					result = tmp
				}
			}
		}
	}
	return result
}

func isPalindrome(s string) bool{

	var isPalindrome bool = true

	for i := 0;i < len(s)/2;i++{
		if s[i] == s[len(s)-i-1]{
			//fmt.Println(s[i],s[len(s)-i-1])
			continue
		}else{
			isPalindrome = false
		}
	}
	return isPalindrome
}

func main() {
	fmt.Println(time.Now())
	fmt.Println(longestPalindrome("vmqjjfnxtyciixhceqyvibhdmivndvxyzzamcrtpywczjmvlodtqbpjayfchpisbiycczpgjdzezzprfyfwiujqbcubohvvyakxfmsyqkysbigwcslofikvurcbjxrccasvyflhwkzlrqowyijfxacvirmyuhtobbpadxvngydlyzudvnyrgnipnpztdyqledweguchivlwfctafeavejkqyxvfqsigjwodxoqeabnhfhuwzgqarehgmhgisqetrhuszoklbywqrtauvsinumhnrmfkbxffkijrbeefjmipocoeddjuemvqqjpzktxecolwzgpdseshzztnvljbntrbkealeemgkapikyleontpwmoltfwfnrtnxcwmvshepsahffekaemmeklzrpmjxjpwqhihkgvnqhysptomfeqsikvnyhnujcgokfddwsqjmqgsqwsggwhxyinfspgukkfowoxaxosmmogxephzhhy"))
	fmt.Println(time.Now().Unix())
	//fmt.Println(isPalindrome("babad"))
}
