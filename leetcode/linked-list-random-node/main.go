package main

import (
    "github.com/davecgh/go-spew/spew"
    "math/rand"
)

type ListNode struct {
    Val  int
    Next *ListNode
}

type Solution struct {
    head *ListNode
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
    return Solution{head:head}
}


/** Returns a random node's value. */
// （No.40）水库抽样问题 https://zhuanlan.zhihu.com/p/46334179
func (this *Solution) GetRandom() int {
    cur := this.head
    ret := cur
    for i:=1 ; cur != nil; i, cur = i+1, cur.Next {
        if rand.Intn(i) == 0 {
            ret = cur
        }
    }
    return ret.Val
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

func main() {
    root := &ListNode{Val:0}
    cur := root
    for i:=1; i < 5; i++ {
        n := &ListNode{Val:i}
        cur.Next = n
        cur = n
    }
    spew.Dump(root)
    rev := reverseList(root)
    spew.Dump(rev)
}
