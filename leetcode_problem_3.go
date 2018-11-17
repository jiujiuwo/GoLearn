package main

import "fmt"

func lengthOfNonRepeatSubStr(s string) int{

	//记录每个字符最后出现的位置
	start := 0
	maxLength := 0
	lastOccured :=make(map[byte]int)
	for i,ch := range []byte(s){

		lastIndex,ok := lastOccured[ch]

		if ok && lastIndex >= start{
			start = lastIndex + 1
		}
		if i - start + 1 > maxLength{
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatSubStr("abcabcbb"))
}
