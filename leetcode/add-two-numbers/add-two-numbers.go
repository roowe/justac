package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {    
    root := &ListNode{}
    iter := root
    carry := 0
    for l1 != nil && l2 != nil {
        carry += l1.Val + l2.Val
        iter.Next = &ListNode{carry%10, nil}
        carry /= 10       
        iter = iter.Next
        l1 = l1.Next
        l2 = l2.Next
    }
    for l1 != nil  {
        carry += l1.Val
        iter.Next = &ListNode{carry%10, nil}
        carry /= 10       
        iter = iter.Next
        l1 = l1.Next
    }
    for l2 != nil  {
        carry += l2.Val
        iter.Next = &ListNode{carry%10, nil}
        carry /= 10       
        iter = iter.Next
        l2 = l2.Next
    }
    if carry > 0 {
        iter.Next = &ListNode{carry, nil}
    }
    return root.Next
}
// 翻转
//post := &ListNode{}
//iter = root.Next
// 7 0 8
// root iter post
//          ->
//          <-
//               ->
//     <-
// root post iter
// for iter.Next != nil {
//     post = iter.Next
//     iter.Next = post.Next
//     post.Next = root.Next
//     root.Next = post
// }
// iter.Next = root
// root = iter.Next.Next
// iter.Next.Next = nil
// for root.Val == 0 {
//     root = root.Next
// }

func parse(v int) *ListNode {
    root := &ListNode{}
    for iter:=root ;v != 0; v/=10 {
        iter.Next = &ListNode{v % 10, nil}
        iter = iter.Next
    }    
    return root.Next
}

func (this *ListNode) String() string{
    s := ""
    for this != nil {
        s += fmt.Sprintf("-> %d %p", this.Val, this.Next)
        this = this.Next
    }
    return s
}

func main() {
    // fmt.Printf("%v\n", parse(342))
    // fmt.Printf("%v\n", parse(465))
    fmt.Printf("%v\n", addTwoNumbers(parse(342), parse(465)))
    fmt.Printf("%v\n", addTwoNumbers(parse(5), parse(5)))
}
