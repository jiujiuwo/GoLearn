package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {

	result := make([][]int,len(nums)*len(nums))
	length := len(nums)
	index :=0

	if(length<=0){
		return result
	}

	for i:=0;i<length;i++{
		for j:=0;j<length;j++{
			for k:=0;k<length;k++{
				if nums[i] + nums[j] + nums[k] == 0 && i !=j && j!=k && i!=k{
					result[index] = make([]int,3)
					result[index][0] = nums[i]
					result[index][1] = nums[j]
					result[index][2] = nums[k]
					if index>0{
						sort.Ints(result[index])
						dup := false
						//为什么是从0到 index+1呢
						for m:=0;m<index;m++{
							sort.Ints(result[m])
							fmt.Println(m,index)
							if result[m][0]==result[index][0]&&result[m][1]==result[index][1]&&result[m][2]==result[index][2]{
								dup = true
								break
							}else{
								continue
							}
						}
						if !dup{
							//fmt.Println(i,j,k)
							index++
						}
					}else{
						index++
					}
				}
			}
		}
	}

	return result[:index]
}

func main() {
	tmp := []int {-1,0,1,2,-1,-4}
	fmt.Println(threeSum(tmp))
}
