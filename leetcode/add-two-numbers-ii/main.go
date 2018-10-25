package main

import "github.com/davecgh/go-spew/spew"

type ListNode struct {
    Val int
    Next *ListNode
}

func addToFront(val int, head *ListNode) *ListNode {
    temp := &ListNode{Val:val}
    temp.Next = head
    return temp
}

func lenList(l *ListNode) int {
    var n int
    for i:=l; i != nil; i = i.Next {
        n++
    }
    return n
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    n1 := lenList(l1)
    n2 := lenList(l2)
    cur1, cur2 := l1, l2
    var res *ListNode
    for ;n1>0 && n2>0; {
        var sum int
        if n1>=n2 {
            sum += cur1.Val
            cur1 = cur1.Next
            n1--
        }
        if n2 > n1 {
            sum += cur2.Val
            cur2 = cur2.Next
            n2--
        }
        // 结果翻转
        res = addToFront(sum, res)
    }
    // 再次翻转
    var ret *ListNode
    var carry int
    for cur := res; cur != nil; cur=cur.Next {
        carry += cur.Val
        ret = addToFront(carry%10, ret)
        carry /= 10
    }
    if carry > 0{
        ret = addToFront(carry%10, ret)
    }
    return ret
}

func main() {
    N1_1 := &ListNode{Val:7}
    N1_2 := &ListNode{Val:2}
    N1_3 := &ListNode{Val:4}
    N1_4 := &ListNode{Val:3}
    N1_1.Next = N1_2
    N1_2.Next = N1_3
    N1_3.Next = N1_4

    N2_1 := &ListNode{Val:5}
    N2_2 := &ListNode{Val:6}
    N2_3 := &ListNode{Val:4}
    N2_1.Next = N2_2
    N2_2.Next = N2_3

    spew.Dump(addTwoNumbers(N1_1, N2_1))
}
