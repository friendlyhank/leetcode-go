package problems

import(
	"testing"
)

/*
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。

示例 1:

输入: 3
输出: "III"
示例 2:

输入: 4
输出: "IV"
示例 3:

输入: 9
输出: "IX"
示例 4:

输入: 58
输出: "LVIII"
解释: L = 50, V = 5, III = 3.
示例 5:

输入: 1994
输出: "MCMXCIV"
解释: M = 1000, CM = 900, XC = 90, IV = 4.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-to-roman
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func intToRoman1(num int) string{
	return FindIntForRoman("",num)
}

func FindIntForRoman(str string,num int)string{
	if num >=1000{
		str += "M"
		return FindIntForRoman(str,num - 1000)
	}else if num >= 900{
		str += "CM"
		return FindIntForRoman(str,num - 900)
	}else if  num >= 500{
		str += "D"
		return FindIntForRoman(str,num - 500)
	}else if num >= 400{
		str += "CD"
		return FindIntForRoman(str,num - 400)
	}else if num >= 100{
		str += "C"
		return FindIntForRoman(str,num - 100)
	}else if num >= 90{
		str += "XC"
		return FindIntForRoman(str,num - 90)
	}else if num >= 50{
		str += "L"
		return FindIntForRoman(str,num - 50)
	}else if num >=40{
		str += "XL"
		return FindIntForRoman(str,num - 40)
	}else if num >= 10{
		str += "X"
		return FindIntForRoman(str,num - 10)
	}else if num >= 9{
		str += "IX"
		return FindIntForRoman(str,num - 9)
	}else if num >= 5{
		str += "V"
		return FindIntForRoman(str,num - 5)
	}else if num >= 4{
		str += "IV"
		return FindIntForRoman(str,num - 4)
	}else if num >= 1{
		str += "I"
		return FindIntForRoman(str,num - 1)
	}

	return str
}

/*
 *自解法(自己写的,不一定最优哦)
 *递归方式检查数字落在哪个区间，然后拼接字符串
 *(卧槽,思路是有，性能还好，怎么能写出这么挫的代码)
 *通过	8 ms	3.3 MB
 */
func TestIntToRoman1(t *testing.T){
	str := intToRoman1(58)
	t.Logf("%v",str)
}
