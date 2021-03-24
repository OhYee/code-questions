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

    cnt := make(map[int]int)

    cur := head
    for cur != nil {
        cnt[cur.Val]++
        cur = cur.Next
    }

    cur = head;
    for cur.Next != nil {
        next := cur.Next
        if cnt[next.Val] > 1 {
            // 需要删除
            cur.Next = next.Next
        } else {
            // 不需要删除
            cur = cur.Next
        }
    }

    if cnt[head.Val] > 1 {
        head = head.Next
    }

    return head
}
