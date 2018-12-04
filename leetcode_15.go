package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	length := len(nums)

	if length<=0{
		return nil
	}
	result := make([][]int,len(nums)*len(nums))
	index :=0

	sort.Ints(nums)

	for i:=0;i<length;i++{
		for j:=i+1;j<length;j++{
			for k:=j+1;k<length;k++{
				if nums[i] + nums[j] + nums[k] == 0{
					result[index] = make([]int,3)
					result[index][0] = nums[i]
					result[index][1] = nums[j]
					result[index][2] = nums[k]
					if index>0{
						//sort.Ints(result[index])
						dup := false
						//为什么是从0到 index+1呢
						for m:=0;m<index;m++{
							//sort.Ints(result[m])
							//fmt.Println(m,index)
							if result[m][0]==result[index][0]&&result[m][1]==result[index][1]&&result[m][2]==result[index][2]{
								dup = true
								break
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
