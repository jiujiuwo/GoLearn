package main

func maxArea(height []int) int {
	var tmp = 0
	var result = 0
	for i:=0;i<len(height);i++{
		for j:=1;j<len(height);j++{

			if height[i]>height[j]{
				tmp = height[j]
			}else{
				tmp = height[i]
			}
			tmp = tmp * (j-i)

			if tmp>result{
				result = tmp
			}
		}
	}
	return result
}

func main() {

}
