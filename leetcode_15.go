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
	result := make([][]int,length*(length-1)*(length-2))
	index :=0

	sort.Ints(nums)
	//fmt.Println(nums)

	for i:=0;i<length;i++{
		//如果没有 nums[i] !=nums[i-1] ,则会有重复 [-4 -1 -1 0 1 2]
		if i==0 || (i >0 && nums[i] !=nums[i-1]){
			low := i+1
			high := length -1
			sum := 0 - nums[i]
			for ;low<high;{
				if nums[low]+nums[high] == sum{
					result[index] = make([]int,3)
					result[index][0] = nums[i]
					result[index][1] = nums[low]
					result[index][2] = nums[high]
					index++
					for;low<high&&nums[low] == nums[low+1];{
						low++
					}
					for;low<high&&nums[high] == nums[high-1];{
						high--
					}
					low++
					high--
				}else if nums[low] + nums[high] < sum{
					low++
				}else{
					high--
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
