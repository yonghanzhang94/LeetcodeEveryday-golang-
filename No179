package main

import (
	"fmt"
	"strconv"
)

func largestNumber(nums []int) string {
	var str string
	for i:=0; i<len(nums); i++{
		for j:=i+1; j<len(nums); j++{
			//fmt.Println(strconv.Itoa(nums[j]))
			if strconv.Itoa(nums[j])+strconv.Itoa(nums[i]) > 
				strconv.Itoa(nums[i])+strconv.Itoa(nums[j]){  // 排序，将最大的数放在前面
				nums[i],nums[j] = nums[j], nums[i]
			}
		}
	}
	if nums[0] == 0{  // 如果第一个数是0，说明整个数组都是0
		return "0"
	}
	for i:=0; i<len(nums); i++{
		//fmt.Println(nums[i])
		str = str + strconv.Itoa(nums[i])
	}
	return str
}

func main() {
	//nums := []int{10,2}
	//nums := []int{3,30,34,5,9}
	//nums := []int{1}
	//nums := []int{10}
	nums := []int{0,0}

	ret := largestNumber(nums)
	fmt.Println(ret)
}
