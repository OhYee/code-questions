/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    var head *ListNode = nil
    var tail *ListNode = nil
    for {
        minIndex := -1
        for i := range lists {
            ptr := lists[i]
            if ptr != nil && (minIndex == -1 || lists[minIndex].Val > ptr.Val) {
                minIndex = i
            }
        }
        if minIndex == -1 {
            break
        }
        ptr := lists[minIndex]
        if head == nil {
            head = ptr
            tail = head
        } else {
            tail.Next = ptr
            tail = ptr
        }
        lists[minIndex] = ptr.Next
    }
    return head
}
