package leetcode

import "fmt"

func longestPalindromeDp(s string) string {

	if len(s) == 0{
		return s
	}

	sLen := len(s)
	result := ""
	max := 0

	dp := make([][]bool,sLen)
	for i:=0; i< sLen;i++{
		dp[i] = make([]bool,sLen)
	}

	//动态规划，i,j分别代表什么？？？
	for i :=0 ;i < sLen; i++{
		for j := 0; j <= i; j++{	//遍历所有子串
			//fmt.Println(result,dp,max,i,j)
			//判断首尾两个字符是否相等
			judge := s[i] == s[j]
			fmt.Println(j,i)
			//动态规划递推式
			//dp[i][j] = (j - i > 2) ? (dp[i+1][j-1] && judge) :judge
			//dp[i][j] 代表什么状态？
			if i - j > 2{		//为什么是 i - j > 2 ?? =2 时，dp[i,j]是不是回文串只和[0:2]内的字符有关
								//当i - j >2时，他是不是回文串才取决于dp[j+1,i-1] && judge
								//i > j+2, i >1,
				dp[j][i] = dp[j+1][i-1] && judge
				//怎么保证[j+1][i-1]的值已经存在？j是小于等于i的，j
			}else{
				dp[j][i] = judge
			}

			if dp[j][i] && ((i - j + 1) > max){
			//	fmt.Println(j,i,max)
				max = i - j + 1
				result = s[j:i+1]
			}
		}
	}
	fmt.Println(dp)
	return result
}

func main() {
	fmt.Println(longestPalindromeDp("aabbaa"))
}
