package main

import "github.com/davecgh/go-spew/spew"

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    fakeHead := &ListNode{Next:head}
    prev := fakeHead
    cur := head
    for cur!=nil {
        start := cur
        end := cur
        i := 0
        for ; end != nil && i<k; i++ {
            end = end.Next
        }
        if i != k {
            break
        }
        for cur = start; cur != end; {
            //fmt.Printf("head %v\n", cur)
            next := cur.Next
            // prev.Next暂存着最后操作的指针
            cur.Next = prev.Next
            prev.Next = cur
            cur = next
        }
        // cur == end
        start.Next = end
        prev = start
    }
    return fakeHead.Next
}

func main() {
    root := &ListNode{Val:1}
    cur := root
    for i:=2; i < 3; i++ {
        n := &ListNode{Val:i}
        cur.Next = n
        cur = n
    }
    spew.Dump(root)
    rev := reverseKGroup(root, 2)
    spew.Dump(rev)
}
