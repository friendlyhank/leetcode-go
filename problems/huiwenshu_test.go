package problems

import (
	"strconv"
	"testing"
)

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
进阶:

你能不将整数转为字符串来解决这个问题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/*
 *自解法(自己写的,不一定最优哦)
 *转化字符串看依次检查字符串是否是对称的，长度是奇数个的时候忽略中间那个
通过	16 ms	5.2 MB
 */
func isPalindrome1(x int) bool {
	str := strconv.Itoa(x)
	for i :=0;i<= len(str) /2;i++{
		if str[i] != str[len(str)-i -1]{
			return false
		}
	}
	return true
}

func TestIsPalindrome1(t *testing.T){
	isPalindrome:= isPalindrome1(1221)
	t.Logf("%v",isPalindrome)
}

/*
 *自解法(自己写的,不一定最优哦)
 *数字反转,然后判断反转之后是否相当即可
通过	20 ms	5 MB
 */
func isPalindrome2(x int) bool {
	//负数和反转之后首位为0
	if  x < 0 || (x!=0 && x % 10 == 0){
		return false
	}

	//反转数字
	var rev int
	tmp := x
	for tmp != 0{
		rev = rev * 10 +tmp % 10
		tmp /= 10
	}

	return x ==  rev
}

func TestIsPalindrome2(t *testing.T){
	isPalindrome:= isPalindrome2(1221)
	t.Logf("%v",isPalindrome)
}
