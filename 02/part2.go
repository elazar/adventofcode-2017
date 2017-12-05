package main

import "bufio"
import "encoding/csv"
import "fmt"
import "os"
import "sort"
import "strconv"

func main() {
    file, _ := os.Open("input.txt")
    reader := csv.NewReader(bufio.NewReader(file))
    reader.Comma = '\t'
    reader.FieldsPerRecord = -1
    rows, _ := reader.ReadAll()
    diffs := make([]int, len(rows))
    var numbers []int
    var first, second int
    for row, columns := range rows {
       numbers = make([]int, len(columns))
       for column, value := range columns {
           numbers[column], _ = strconv.Atoi(value)
       }
       sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
       for first = 0 ; first < len(columns) - 1 ; first++ {
           for second = first + 1; second < len(columns); second++ {
                if numbers[first] % numbers[second] == 0 {
                    diffs[row] = numbers[first] / numbers[second]
                    // fmt.Printf("%d %d %d\n", numbers[first], numbers[second], diffs[row])
                }
           }
       }
    }
    sum := 0
    for _, diff := range diffs {
        sum += diff
    }
    fmt.Println(sum)
}
