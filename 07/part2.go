package main

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"

type Program struct {
    own_weight int
    total_weight int
    children []string
    parent string
}

func CalculateTotalWeights(programs map[string]*Program, program string) int {
    total_weight := programs[program].own_weight
    for _, child := range programs[program].children {
        total_weight += CalculateTotalWeights(programs, child)
    }
    programs[program].total_weight = total_weight
    return total_weight
}

func OutputTree(programs map[string]*Program, program string, indent int, limit int) {
    fmt.Printf(
        "%s%s (%d, %d)\n",
        strings.Repeat(" ", indent*2),
        program,
        programs[program].own_weight,
        programs[program].total_weight,
    )
    if limit == 0 || indent < limit {
        for _, child := range programs[program].children {
            OutputTree(programs, child, indent + 1, limit)
        }
    }
}

func FindCorrectWeight(programs map[string]*Program, program string) int {
    weights := make(map[int]int)
    var child_weight int
    var children []string = programs[program].children
    for _, child := range children {
        child_weight = programs[child].total_weight
        _, present := weights[child_weight]
        if !present {
            weights[child_weight] = 0
        }
        weights[child_weight]++
    }
    if len(weights) == 1 {
        parent := programs[program].parent
        own_weight := programs[program].own_weight
        lone_weight := programs[program].total_weight
        var other_weight int
        for _, child := range programs[parent].children {
            if child != parent {
                other_weight = programs[child].total_weight
                break
            }
        }
        return own_weight - (lone_weight - other_weight)
    } else {
        var lone_weight int
        var child string
        for child_weight, count := range weights {
            if count == 1 {
                lone_weight = child_weight
                break
            }
        }
        for _, child = range children {
            if programs[child].total_weight == lone_weight {
                return FindCorrectWeight(programs, child)
            }
        }
    }
    return 0
}

func main() {
    file, _ := os.Open("input.txt")
    reader := bufio.NewReader(file)
    var columns []string
    var name string
    var own_weight int
    var children []string
    programs := make(map[string]*Program)

    for {
        line, _, err := reader.ReadLine()
        if err != nil {
            break
        }
        columns = strings.SplitN(string(line), " ", 4)
        name = columns[0]
        own_weight, _ = strconv.Atoi(strings.Trim(columns[1], "()"))
        if len(columns) > 2 {
            children = strings.Split(columns[3], ", ")
        } else {
            children = make([]string, 0)
        }
        programs[name] = &Program{own_weight: own_weight, children: children}
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

    CalculateTotalWeights(programs, root)

    fmt.Println(FindCorrectWeight(programs, root))
}
