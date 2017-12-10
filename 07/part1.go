package main

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"

type Program struct {
    weight int
    children []string
    parent string
}

func main() {
    file, _ := os.Open("input.txt")
    reader := bufio.NewReader(file)
    var columns []string
    var name string
    var weight int
    var children []string
    programs := make(map[string]*Program)

    for {
        line, _, err := reader.ReadLine()
        if err != nil {
            break
        }
        columns = strings.SplitN(string(line), " ", 4)
        name = columns[0]
        weight, _ = strconv.Atoi(strings.Trim(columns[1], "()"))
        if len(columns) > 2 {
            children = strings.Split(columns[3], ", ")
        } else {
            children = make([]string, 0)
        }
        programs[name] = &Program{weight: weight, children: children}
    }

    for parent, program := range programs {
        for _, child := range program.children {
            programs[child].parent = parent
        }
    }

    var root string
    for name, program := range programs {
        if len(program.parent) == 0 {
            root = name
            break
        }
    }
    fmt.Println(root)
}
