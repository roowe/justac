package main
 
import "fmt"
 
type ListNode struct {
	Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var fast *ListNode = head
    for i := 0; i<n; i++ {
        fast = fast.Next        
    }
    if fast == nil {
        return head.Next
    }
    var cur *ListNode = head
    for fast.Next != nil {
        fast = fast.Next
        cur = cur.Next
    }
    cur.Next = cur.Next.Next
    return head
}

func p(head *ListNode) {
    for ;head != nil;head=head.Next {
        fmt.Println(head.Val, head.Next)
    }
}

func main() {
    list := make([]ListNode, 100)
    for i := 0; i<10; i++ {
        v := &list[i]
        v.Val = i
        v.Next = &list[i+1]
    }
    // fmt.Printf("len=%d cap=%d %v\n", len(list), cap(list), list)
    var head *ListNode = &list[0]
    p(head)
    var rest *ListNode = removeNthFromEnd(head, 5)
    p(rest)
}
