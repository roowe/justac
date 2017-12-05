package main

import "fmt"

type ListNode struct {
   Val int
   Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := new(ListNode)
    ptr := head
    for ;l1 != nil && l2 != nil;{
        if l1.Val < l2.Val {
            ptr.Next = l1
            ptr = l1
            l1 = l1.Next
        } else {
            ptr.Next = l2
            ptr = l2
            l2 = l2.Next
        }
    }
    if l1 != nil {
        ptr.Next = l1        
    }
    if l2 != nil {
        ptr.Next = l2        
    }
    return head.Next
}

func p(head *ListNode) {
    fmt.Println("------start--------------")
    for ;head != nil;head=head.Next {
        fmt.Println(head.Val, head.Next)
    }
    fmt.Println("------end--------------")
}

func main() {
   list := make([]ListNode, 100)
   for i := 0; i<10; i++ {
        v := &list[i]
        v.Val = i
        v.Next = &list[i+1]
   }
   list[9].Next = nil
   for i := 20; i<30; i++ {
        v := &list[i]
        v.Val = i-25
        v.Next = &list[i+1]
    }
    list[29].Next = nil
    // fmt.Printf("len=%d cap=%d %v\n", len(list), cap(list), list)
    var head1 *ListNode = &list[0]
    var head2 *ListNode = &list[20]
    p(head1)
    p(head2)
    var rest *ListNode = mergeTwoLists(head1, head2)
    p(rest)
}
