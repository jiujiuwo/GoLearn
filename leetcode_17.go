package main

import "fmt"

func combine(result []string,runes []string) []string{
	//fmt.Println(result)
	//fmt.Println(runes)
	//对一开始，第一个字符进行特殊处理
	if len(result)==0{
		for j:=0;j<len(runes);j++{
			result = append(result,runes[j])
		}
	}else{//从第二个字符开始，依次为之前的字符，加上当前字符数组的每个元素，append到result中
		lenResult := len(result)
		for i:=0; i<lenResult; i++{
			for j:=0;j<len(runes);j++{
				result = append(result,result[i]+runes[j])
			}
		}
	}
	//fmt.Println(result)
	return result
}

func letterCombinations(digits string) []string {
	//用来存储结果
	result := []string{}

	if(len(digits)==0){
		return result
	}

	//首先，将数字和字母的映射关系写好
	maps := map[rune][]string{
		'0':{},
		'2':{"a","b","c"},
		'3':{"d","e","f"},
		'4':{"g","h","i"},
		'5':{"j","k","l"},
		'6':{"m","n","o"},
		'7':{"p","q","r","s"},
		'8':{"t","u","v"},
		'9':{"w","x","y","z"},
	}

	//计算真正的结果中包含的元素个数，即，字符串长度为len(digits)的个数
	index := 1
	for _,digit := range []rune(digits){
		//fmt.Println(index,maps[ditit])
		result = combine(result,maps[digit])
		index = index * len(maps[digit])
	}
	//fmt.Println("over")
	//fmt.Println(maps)
	//fmt.Println(result)
	//从结果数组去，除去前面长度不足的元素
	return result[len(result)-index:]
}

func main() {
	fmt.Println(letterCombinations("234"))
}
