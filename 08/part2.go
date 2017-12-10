package main

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"

type Instruction struct {
    register string
    instruction string
    quantity int
    condition_register string
    condition_operator string
    condition_quantity int
}

func main() {
    file, _ := os.Open("input.txt")
    reader := bufio.NewReader(file)
    var columns []string
    var instruction Instruction
    var condition bool
    var max int
    registers := make(map[string]int)
    for {
        line, _, err := reader.ReadLine()
        if err != nil {
            break
        }
        columns = strings.Split(string(line), " ")
        instruction = Instruction{
            register: columns[0],
            instruction: columns[1],
            condition_register: columns[4],
            condition_operator: columns[5],
        }
        instruction.quantity, _ = strconv.Atoi(columns[2])
        instruction.condition_quantity, _ = strconv.Atoi(columns[6])
        _, present := registers[instruction.condition_register]
        if !present {
            registers[instruction.condition_register] = 0
        }
        register := registers[instruction.condition_register]
        quantity := instruction.condition_quantity
        switch instruction.condition_operator {
        case "!=":
            condition = register != quantity
        case "==":
            condition = register == quantity
        case ">=":
            condition = register >= quantity
        case "<=":
            condition = register <= quantity
        case ">":
            condition = register > quantity
        case "<":
            condition = register < quantity
        }
        if condition == true {
            if instruction.instruction == "inc" {
                registers[instruction.register] += instruction.quantity
            } else if instruction.instruction == "dec" {
                registers[instruction.register] -= instruction.quantity
            }
        }
        if registers[instruction.register] > max {
            max = registers[instruction.register]
        }
    }
    fmt.Printf("Overall: %d\n", max)
    max = 0
    for _, value := range registers {
        if value > max {
            max = value
        }
    }
    fmt.Printf("Final: %d\n", max)
}
