package main

import "bufio"
import "fmt"
import "os"
import "strings"

func isValid(line []byte) bool {
    var words []string
    var wordmap map[string]bool
    var w int
    var present bool
    words = strings.Split(string(line), " ")
    wordmap = make(map[string]bool, 0)
    //fmt.Print(string(line))
    for w = 0 ; w < len(words) ; w++ {
        _, present = wordmap[words[w]]
        if present {
            //fmt.Printf(" - INVALID #%d", invalid)
            return false
        }
        wordmap[words[w]] = true
    }
    return true
}

func main() {
    file, _ := os.Open("input.txt")
    reader := bufio.NewReader(file)
    lines := 0
    invalid := 0
    for {
        line, _, err := reader.ReadLine()
        if err != nil {
            break
        }
        lines++
        if !isValid(line) {
            invalid++
        }
    }
    fmt.Println(lines - invalid)
}
