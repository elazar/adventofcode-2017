package main

import "bufio"
import "encoding/csv"
import "fmt"
import "os"
import "strconv"

func max(numbers []int) int {
    max := 0
    for _, number := range numbers {
        if (number > max) {
            max = number
        }
    }
    return max
}

func min(numbers []int) int {
    min := int(^uint(0) >> 1)
    for _, number := range numbers {
        if (number < min) {
            min = number
        }
    }
    return min
}

func main() {
    file, _ := os.Open("input.txt")
    reader := csv.NewReader(bufio.NewReader(file))
    reader.Comma = '\t'
    reader.FieldsPerRecord = -1
    rows, _ := reader.ReadAll()
    diffs := make([]int, len(rows))
    var numbers []int
    for row, columns := range rows {
       numbers = make([]int, len(columns))
       for column, value := range columns {
           numbers[column], _ = strconv.Atoi(value)
       }
       diffs[row] = max(numbers) - min(numbers)
       fmt.Printf("%d %d %d\n", max(numbers), min(numbers), diffs[row])
    }
    sum := 0
    for _, diff := range diffs {
        sum += diff
    }
    fmt.Println(sum)
}
