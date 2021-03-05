/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    tail := head
    ptr := tail
    first := true
    for {
        newhead, newtail := reverse(ptr, k)
        debug(newhead, newtail)
        if newhead == nil {
            break
        }

        if first {
            head = newhead
            first = false
        } else {
            tail.Next = newhead
        }
        tail = newtail
        ptr = newtail.Next
    }
    return head
}

func debug(f, t*ListNode ) {
    if f == nil || f == t {
        if f != nil {
            fmt.Printf("%d", f.Val)
        } else {
            fmt.Printf("nil")
        }
        fmt.Printf("\n")
        return
    }
    fmt.Printf("%d -> ",f.Val)
    debug(f.Next, t)
}

func reverse(head *ListNode, k int) (*ListNode, *ListNode) {
    newhead := head
    ptr := head
    for i := 0; i < k; i++ {
        if ptr == nil {
            return nil, nil
        }
        ptr = ptr.Next
    }

    ptr = head.Next
    for i := 0; i < k - 1; i++ {
        temp := ptr.Next
        ptr.Next = newhead
        head.Next = temp
        newhead = ptr
        ptr = temp
    }

    return newhead, head
}
