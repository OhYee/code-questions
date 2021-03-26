/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    m := make(map[int]struct{})

    m[head.Val] = struct{}{}

    cur := head 
    for cur.Next != nil {
        if _, ex := m[cur.Next.Val]; ex {
            cur.Next = cur.Next.Next
        } else {
            m[cur.Next.Val] = struct{}{}
            cur = cur.Next
        }
    }

    return head
}
