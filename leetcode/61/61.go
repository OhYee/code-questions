/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }

    it := head
    l := 0
    for it!= nil {
        it = it.Next
        l++
    }
    k %= l
    
        it1 := head
        it2 := head
        count := 0
        for it2.Next != nil {
            if count >= k {
                it1 = it1.Next
            }
            it2 = it2.Next
            count++
        }

        it2.Next = head
        head = it1.Next
        it1.Next = nil
  
    return head
}
