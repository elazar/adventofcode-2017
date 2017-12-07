package main

import "fmt"

func IntArraysEqual(n []int, m []int) bool {
    if len(n) != len(m) {
        return false
    }
    var index int
    for index = 0 ; index < len(n) ; index++ {
        if n[index] != m[index] {
            return false
        }
    }
    return true
}

func main() {
    input := []int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}
    seen := make([][]int, 0)
    max_index := 0
    var max, index int
    found := false
    for !found {
        seen = append(seen, make([]int, len(input)))
        copy(seen[len(seen)-1], input)
        for index = 0 ; index < len(input); index++ {
            if input[index] > max {
                max = input[index]
                max_index = index
            }
        }
        //fmt.Printf("input %+v, max %d, max_index %d\n", input, max, max_index)
        input[max_index] = 0
        for index = (max_index + 1) % len(input) ; max > 0 ; index = (index + 1) % len(input) {
            input[index]++
            max--
        }
        for index = 0 ; index < len(seen); index++ {
            if IntArraysEqual(input, seen[index]) {
                //fmt.Printf("duplicate %+v\n", input)
                found = true
                break
            }
        }
    }
    fmt.Println(len(seen))
}
