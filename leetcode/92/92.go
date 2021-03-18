/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func reverse(head *ListNode, k int) (newHead *ListNode, newTail *ListNode){
    cur := head.Next
    newHead = head
    newTail = head
    for i:=1; i<k; i++ {
        head.Next = cur.Next
        cur.Next = newHead
        newHead = cur
        cur = head.Next
    }
    return
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
    cur := head
    for i:=0; i<left-2; i++ {
        cur = cur.Next
    }
    if left < 2 {
        head, _ = reverse(cur, right-left+1)
    } else {
        cur.Next, _ = reverse(cur.Next, right-left+1)

    }
    return head
}
