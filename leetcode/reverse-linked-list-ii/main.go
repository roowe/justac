package main

import (
    "github.com/davecgh/go-spew/spew"
)

type ListNode struct {
    Val  int
    Next *ListNode
}
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    // node_1...node_m
    //            m-1,m
    //            pre,end,pre_next,...,end,...cur...node_n
    prev := &ListNode{
        Next: head,
    }
    fakeHead := prev
    i := 1
    for ;i<m; i, prev = i+1, prev.Next{
        //fmt.Printf("prev %+v\n", prev)
    }
    if prev.Next == nil {
        return head
    }
    end := prev.Next
    cur := end.Next

    for ; i<n; i++{
        //fmt.Printf("head %v\n", cur)
        next := cur.Next
        // prev.Next暂存着最后操作的指针
        cur.Next = prev.Next
        prev.Next = cur
        cur = next
    }
    end.Next = cur
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
    rev := reverseBetween(root, 1, 2)
    spew.Dump(rev)
    spew.Dump(reverseBetween(&ListNode{Val:1}, 1, 1))

}
