package main

import "fmt"
import "os"
import "strconv"

func main() {
    input := os.Args[1]
    length := len(input)
    first := 0
    sum := 0
    var second int
    var current int
    for {
        second = (first + 1) % length
        if input[first] == input[second] {
            current, _ = strconv.Atoi(string(input[first]))
            sum += current
        }
        //fmt.Printf("%d %d %s %s %d\n", first, second, string(input[first]), string(input[second]), sum);
        if second == 0 {
            break;
        }
        first = second
    }
    fmt.Println(sum);
}
