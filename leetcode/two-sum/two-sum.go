package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := make(map[int] int)	
	for i1, v1 := range nums {
		v2 := target - v1
		i2, ok := hash[v2]
		if ok {
			return []int{i2, i1}
		} else {
			hash[v1] = i1
		}
	}
	return []int{};
}
func twoSum0(nums []int, target int) []int {
	ret := []int{}
	len := len(nums)
	for i:=0; i<len; i++ {
		for j:=i+1; j<len; j++ {
			if nums[i]+nums[j] == target {
				ret = []int{i,j}
			}
		}
	}
	return ret;
}
func main() {
	data := []int{2, 7, 11, 15}
	ret := twoSum(data, 9)
	fmt.Printf("%v\n", ret)
}