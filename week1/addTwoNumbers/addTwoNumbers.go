package addTwoNumbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode-cn.com/problems/add-two-numbers/submissions/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{
		0,
		nil,
	}
	cursorNode := result
	//进位
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		l1Val := 0
		if l1 != nil {
			l1Val = l1.Val
		}
		l2Val := 0
		if l2 != nil {
			l2Val = l2.Val
		}

		sumVal := l1Val + l2Val + carry
		carry = sumVal / 10

		tempNode := ListNode{
			sumVal % 10,
			nil,
		}

		cursorNode.Next = &tempNode
		cursorNode = &tempNode

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return result.Next
}
