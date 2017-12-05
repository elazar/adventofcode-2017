package main

import "bufio"
import "fmt"
import "os"
import "strings"

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    length := len(s)

    var smap = make(map[byte]int, 0)
    var tmap = make(map[byte]int, 0)
    var i int
    var present bool
    for i = 0 ; i < length ; i++ {
        _, present = smap[s[i]]
        if !present {
            smap[s[i]] = 0
        }
        smap[s[i]]++;

        _, present = tmap[t[i]]
        if !present {
            tmap[t[i]] = 0
        }
        tmap[t[i]]++;
    }

    for b, c := range smap {
        if tmap[b] != c {
            return false
        }
    }
    return true
}

func isValid(line []byte) bool {
    var words []string
    var w, y int
    words = strings.Split(string(line), " ")
    //fmt.Print(string(line))
    for w = 0 ; w < len(words) - 1 ; w++ {
        for y = w + 1 ; y < len(words) ; y++ {
            if isAnagram(words[w], words[y]) {
                return false
            }
        }
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
