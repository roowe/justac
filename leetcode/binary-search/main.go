package main

func search(nums []int, target int) int {
    low, high := 0, len(nums)-1
    for low <= high {
        m := (low+high)/2
        if target > nums[m] {
            // 右边
            low = m+1
        } else if target < nums[m] {
            // 左边
            high = m-1
        } else {
            return m
        }
    }
    return -1
}

func main() {

}
