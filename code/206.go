/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    lst := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return lst
}

// Follow up: A linked list can be reversed either iteratively or recursively.
func reverseList2(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    res := &ListNode{head.Val, nil}
    for head = head.Next; head != nil; head = head.Next {
        cur := &ListNode{head.Val, res}
        res = cur
    }
    return res
}

