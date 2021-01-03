/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */


struct ListNode* partition(struct ListNode* head, int x){
    struct ListNode *head_a = NULL, *a = NULL, *head_b = NULL, *b = NULL, *t = head;
    while (t != NULL) {
        if (t->val < x) {
            if (a == NULL) {
                a = t;
                head_a = a;
            } else {
                a->next = t;
                a = t;
            }
        } else {
            if (b == NULL) {
                b = t;
                head_b = b;
            } else {
                b->next = t;
                b = t;
            }
        }
        struct ListNode *next = t->next;
        t->next = NULL;
        t = next;
    }
    if (a == NULL) {
        return head_b;
    } else {
        a->next = head_b;
        return head_a;
    }
}