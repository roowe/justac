package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    l1 := len(nums1)
    l2 := len(nums2)
    if (l1+l2)%2 == 1 {
        return float64(findK(nums1, nums2, 1+(l1+l2)/2))
    } else {
        return float64(findK(nums1, nums2, (l1+l2)/2) + findK(nums1, nums2, 1+(l1+l2)/2))/2.0
    }
}

func findK(nums1 []int, nums2 []int, k int) int {
    b1, e1 := 0, len(nums1)-1
    b2, e2 := 0, len(nums2)-1
    for b1 <= e1 && b2 <= e2 {
        m1 := (b1+e1)/2
        m2 := (b2+e2)/2
        halfL := m1+m2-b1-b2+2
        if (nums1[m1] < nums2[m2]) {
            if (halfL > k) {
                e2 = m2-1
            } else {
                k-=(m1-b1+1)
                b1 = m1+1
            }
        } else {
            if (halfL > k) {
                e1 = m1-1
            } else {
                k-=(m2-b2+1)
                b2 = m2+1
            }
        }
    }
    if (b1 > e1) {
        return nums2[b2+k-1]
    }
    return nums1[b1+k-1]
}

// func findK2(nums1 []int, b1 int, e1 int, nums2 []int, b2 int, e2 int, k int) int {
//     if (b1 > e1) {
//         return nums2[b2+k-1]
//     }
//     if (b2 > e2) {
//         return nums1[b1+k-1]
//     }
//     m1 := (b1+e1)/2
//     m2 := (b2+e2)/2
//     halfL := m1+m2-b1-b2+2
//     if (nums1[m1] < nums2[m2]) {
//         if (halfL > k) {
//             return findK(nums1, b1, e1, nums2, b2, m2-1, k)
//         } else {
//             return findK(nums1, m1+1, e1, nums2, b2, e2, k-(m1-b1+1))
//         }
//     } else {
//         if (halfL > k) {
//             return findK(nums1, b1, m1-1, nums2, b2, e2, k)
//         } else {
//             return findK(nums1, b1, e1, nums2, m2+1, e2, k-(m2-b2+1))
//         }
//     }
// }
func main() {
    fmt.Printf("%v\n", findMedianSortedArrays(
        []int{1, 3}, []int{2}))
    fmt.Printf("%v\n", findMedianSortedArrays(
        []int{1, 2}, []int{3, 4}))
}
