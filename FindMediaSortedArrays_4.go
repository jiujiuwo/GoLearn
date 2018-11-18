package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	index := (len(nums1)+len(nums2)) / 2
	//fmt.Println(index)

	var median float64 = 0

	newArray := make([]int,index+1)

	m,n := 0,0

	for i := 0;i <=index;i++{
		if m < len(nums1) && n < len(nums2){
			if nums1[m] > nums2[n]{
				newArray[i] = nums2[n]
				n++
			}else{
				newArray[i] = nums1[m]
				m++
			}
			continue
		}

		if(m < len(nums1)){
			newArray[i] = nums1[m]
			m++
		}
		if(n < len(nums2)){
			newArray[i] = nums2[n]
			n++
		}
	}

	//fmt.Println(newArray)

	if (len(nums2) + len(nums1)) % 2 == 0{
		median =  (float64(newArray[index] + newArray[index -1]) / 2)
	}else{
		median =  float64(newArray[index])
	}
	return median
}

func main() {
	nums1 := []int{1,2}
	nums2 := []int{3,4}
	fmt.Println(findMedianSortedArrays(nums1,nums2))
}
