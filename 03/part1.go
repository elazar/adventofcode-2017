package main

import "fmt"
import "math"
import "os"
import "strconv"

func main() {
    input, _ := strconv.Atoi(os.Args[1])
    horizontal := 0
    vertical := 0
    cell := 1
    offset := 0
    direction := 1
    var h, v int
    for cell < input {
        offset++
        for h = 0 ; cell < input && h < offset; h++ {
            cell++
            horizontal += direction
            fmt.Printf("o %d, c %d, h %d, v %d, d %d\n", offset, cell, horizontal, vertical, direction)
        }
        for v = 0; cell < input && v < offset ; v++ {
            cell++
            vertical += direction
            fmt.Printf("o %d, c %d, h %d, v %d, d %d\n", offset, cell, horizontal, vertical, direction)
        }
        direction *= -1
    }
    steps := math.Abs(float64(horizontal)) + math.Abs(float64(vertical))
    fmt.Println(steps)
}
