package main

import "github.com/davecgh/go-spew/spew"

type ListNode struct {
    Val  int
    Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
    oddH := new(ListNode)
    evenH := new(ListNode)
    oddCur := oddH
    evenCur := evenH
    cur := head

    for cur != nil {
        // 1
        oddCur.Next = cur

        // 2
        cur = cur.Next
        evenCur.Next = cur
        oddCur.Next.Next = nil
        oddCur = oddCur.Next

        if cur == nil {
            break
        }
        // 3
        cur = cur.Next
        evenCur = evenCur.Next
    }
    oddCur.Next = evenH.Next
    return oddH.Next
}

func main() {
    root := &ListNode{Val:1}
    cur := root
    for i:=2; i < 10; i++ {
        n := &ListNode{Val:i}
        cur.Next = n
        cur = n
    }
    spew.Dump(root)
    rev := oddEvenList(root)
    spew.Dump(rev)
}