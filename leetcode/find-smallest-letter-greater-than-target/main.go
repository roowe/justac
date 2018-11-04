package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
    low, high := 0, len(letters)-1
    for low <= high {
        m := (low + high) / 2
        if letters[m] <= target {
            low = m + 1
        } else {
            if m == 0 || target >= letters[m-1] {
                // find
                return letters[m]
            } else {
                high = m - 1
            }
        }
    }
    return letters[0]
}

func main() {
    fmt.Printf("%c\n", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
    fmt.Printf("%c\n", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))
    fmt.Printf("%c\n", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd'))
    fmt.Printf("%c\n", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'g'))
    fmt.Printf("%c\n", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'j'))
    fmt.Printf("%c\n", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'k'))
}
