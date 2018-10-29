package main

import (
    "fmt"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func rob(root *TreeNode) int {
    v0, v1 := dfs2(root)
    return max(v0, v1)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func dfs2(root *TreeNode) (int, int) {
    if root == nil {
        return 0, 0
    }
    //spew.Dump(root)
    l0, l1 := dfs2(root.Left)
    r0, r1 := dfs2(root.Right)
    v0 := root.Val + l0 + r0
    v1 := max(l0, l1) + max(r0, r1)
    //spew.Println(root.Val, v0, v1)
    return v1, v0
}

func dfs(root *TreeNode) int {
    if root == nil {
        return 0
    }
    var sum1 = root.Val
    if root.Left != nil {
        sum1 += dfs(root.Left.Left) + dfs(root.Left.Right)
    }

    if root.Right != nil {
        sum1 += dfs(root.Right.Left) + dfs(root.Right.Right)
    }
    var sum2 = dfs(root.Left) + dfs(root.Right)
    if sum1 > sum2 {
        return sum1
    }
    return sum2
}

func main() {
    fmt.Println(rob(Test1()))
    fmt.Println(rob(Test2()))
}

func Test3() *TreeNode {
    root := TreeNode{Val: 3}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 1}
    return &root
}
func Test1() *TreeNode {
    root := TreeNode{Val: 3}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}

    root.Left.Right = &TreeNode{Val: 3}
    root.Right.Right = &TreeNode{Val: 1}
    return &root
}

func Test2() *TreeNode {
    root := TreeNode{Val: 3}
    root.Left = &TreeNode{Val: 4}
    root.Right = &TreeNode{Val: 5}

    root.Left.Left = &TreeNode{Val: 1}
    root.Left.Right = &TreeNode{Val: 3}
    root.Right.Right = &TreeNode{Val: 1}
    return &root
}
