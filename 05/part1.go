package main

import "bufio"
import "fmt"
import "os"
import "strconv"

func main() {
    file, _ := os.Open("input.txt")
    reader := bufio.NewReader(file)
	var offsets = make([]int, 0)
	var offset int
    for {
        line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		offset, _ = strconv.Atoi(string(line))
		offsets = append(offsets, offset)
	}
	current := 0
	steps := 0
	var jump int
	for current < len(offsets) {
		steps++
		jump = offsets[current]
		offsets[current]++
		current += jump
	}
	fmt.Println(steps)
}
