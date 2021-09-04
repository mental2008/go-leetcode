/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    l3 := &ListNode{0, nil}
    pt, add := l3, 0
    for {
        res := add
        add = 0
        if l1 != nil {
            res += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            res += l2.Val
            l2 = l2.Next
        }
        if res >= 10 {
            res = res % 10
            add = 1
        }
        l3.Val = res
        if l1 == nil && l2 == nil {
            break
        }
        l3.Next = &ListNode{0, nil}
        l3 = l3.Next
    }
    if add > 0 {
        l3.Next = &ListNode{1, nil}
    }
    return pt
}
