package main

import (
	"fmt"
	"math"
)

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i:=0;i<len(nums);i++{
		for j:=i+1; j<=i+k && j<len(nums); j++ {  // 只找后续k个数，就可以满足第二个条件
			if math.Abs(float64(nums[i]-nums[j]))<=float64(t){
				print(i, j)
				return true
			}
		}
	}
	return false
}

func main() {
	//nums := []int{1,2,3,1}
	//k := 3
	//t := 0  //true

	//nums := []int{1,0,1,1}
	//k := 1
	//t := 2  //true

	nums := []int{1,5,9,1,5,9}
	k := 2
	t := 3  //false

	ret := containsNearbyAlmostDuplicate(nums, k, t)
	fmt.Println(ret)
}
