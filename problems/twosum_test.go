package problems

import (
	"fmt"
	"testing"
)

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
  暴力解法
  时间复杂度n2
  通过	60 ms	3 MB	Golang
 */
func twoSum1(nums []int, target int) []int {
	for i:=0;i<len(nums);i++{
		for j := i+1;j<len(nums);j++{
			if  nums[j] == target - nums[i]{
				return  []int{i,j}
			}
		}
	}
	return []int{0,0}
}

func TestTwosum1(t *testing.T){
	var nums =[]int{3,2,3}
	var target=6
	keys := twoSum1(nums,target)

	fmt.Println(keys)
}

/**
	比较常用到且比较经典的一遍Hash用法
	时间复杂度n
	通过	4 ms	3.8 MB
 */
func twoSum2(nums []int, target int) []int {
	var mapnum = make(map[int]int,0)
	for i,v := range nums{
		j,exit := mapnum[v]
		if exit{
			return []int{i,j}
		}
		mapnum[target-v]=i
	}
	return nil
}

func TestTwosum2(t *testing.T){
	var nums =[]int{3,2,3}
	var target=6
	keys := twoSum2(nums,target)

	fmt.Println(keys)
}
