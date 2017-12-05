package main

import "fmt"
import "os"
import "strconv"

func main() {
    input := os.Args[1]
    length := len(input)
    half := length / 2
    first := 0
    sum := 0
    var second int
    var current int
    for {
        second = (first + half) % length
        if input[first] == input[second] {
            current, _ = strconv.Atoi(string(input[first]))
            sum += current
        }
        //fmt.Printf("%d %d %s %s %d\n", first, second, string(input[first]), string(input[second]), sum);
        if first == length - 1 {
            break;
        }
        first = first + 1
    }
    fmt.Println(sum);
}
