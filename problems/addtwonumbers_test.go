package problems


import(
	"testing"
)

/*
两数相加

给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

//辅助方法slice转ListNode
func ConverSliceToNode(nums []int) *ListNode{
	var listNode *ListNode
	var temp *ListNode
	for _,num := range nums{
		if temp == nil{
			temp = &ListNode{Val:num}
			listNode = temp
		}else{
			temp.Next = &ListNode{Val:num}
			temp = temp.Next
		}
	}
	return listNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var num3 int64
	var i int64=1
	for l1 != nil || l2 != nil{

		if l1 != nil{
			num3 += int64(l1.Val)*i
			l1 = l1.Next
		}

		if l2 != nil{
			num3 += int64(l2.Val)*i
			l2 = l2.Next
		}

		i *= 10
	}

	if num3 == 0{
		return &ListNode{Val:0}
	}

	var l3 *ListNode
	var temp *ListNode
	for num3 != 0{
		currentNode :=&ListNode{Val:int(num3 % 10)}
		num3 /= 10

		if l3 == nil{
			l3 = currentNode
			temp = l3
		}else{
			temp.Next = currentNode
			temp =temp.Next
		}
	}
	return l3
}

/*
 *写了一个比较挫的实现
 *常规来说没问题，但是通不过测试用例,因为这样方式数字太大会溢出
 *题目是想通过两数由表头开始一次相加的方式实现
 */
func TestAddTwoNumbers(t *testing.T){
	sliceOne := []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}
	sliceOne2 := []int{5,6,4}

	l1 :=ConverSliceToNode(sliceOne)
	l2 := ConverSliceToNode(sliceOne2)

	l3 := addTwoNumbers(l1,l2)

	t.Logf("%v",l3)
}
