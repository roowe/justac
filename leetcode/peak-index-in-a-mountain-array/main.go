package main

import "fmt"

func peakIndexInMountainArray(A []int) int {
    n := len(A)
    low, high := 0, n-1
    for low <= high {
        mid := (low+high)/2
        if A[mid] > A[mid-1] &&
            A[mid] > A[mid+1] {
               return mid
        } else if A[mid] < A[mid+1] {
            low = mid+1
        } else {
            high = mid - 1
        }
    }
    return -1
}

func main()  {
    fmt.Println(peakIndexInMountainArray([]int{0,1,0}))
    fmt.Println(peakIndexInMountainArray([]int{0,2,1,0}))
}