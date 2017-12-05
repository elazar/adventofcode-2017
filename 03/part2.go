package main

import "fmt"
import "math"
import "os"
import "strconv"

type Cell struct {
    number int
    horizontal int
    vertical int
    offset int
    direction int
    value int
    adjacencies []*Cell
}

func IntAbs(n int, m int) float64 {
    return math.Abs(float64(n) - float64(m));
}

func (c Cell) IsAdjacentTo(o Cell) bool {
    return o.number < c.number &&
        (IntAbs(o.vertical, c.vertical) <= 1 &&
        IntAbs(o.horizontal, c.horizontal) <= 1)
}

func AddCell(
    number int,
    horizontal int,
    vertical int,
    offset int,
    direction int,
    cells []Cell,
) []Cell {
    var value int = 0
    if number == 1 {
        value = 1
    }
    cell := Cell{
        number: number,
        horizontal: horizontal,
        vertical: vertical,
        offset: offset,
        direction: direction,
        value: value,
        adjacencies: make([]*Cell, 0),
    }

    var c int
    for c = len(cells) - 1 ; c >= 0 ; c-- {
        //fmt.Printf("c %+v o %+v\n", cell, cells[c])
        if cell.IsAdjacentTo(cells[c]) {
            cell.adjacencies = append(cell.adjacencies, &cells[c])
            cell.value += cells[c].value
        }
    }

    return append(cells, cell)
}

func main() {
    input, _ := strconv.Atoi(os.Args[1])
    horizontal := 0
    vertical := 0
    number := 1
    offset := 0
    direction := 1
    lastCellValue := 0
    var h, v int

    cells := make([]Cell, 0)
    cells = AddCell(
        number,
        horizontal,
        vertical,
        offset,
        direction,
        cells,
    )

    for lastCellValue <= input {
        offset++
        for h = 0 ; lastCellValue <= input && h < offset; h++ {
            number++
            horizontal += direction
            cells = AddCell(
                number,
                horizontal,
                vertical,
                offset,
                direction,
                cells,
            )
            lastCellValue = cells[len(cells)-1].value
            //fmt.Printf("%+v\n", cells[len(cells)-1])
        }
        for v = 0; lastCellValue <= input && v < offset ; v++ {
            number++
            vertical += direction
            cells = AddCell(
                number,
                horizontal,
                vertical,
                offset,
                direction,
                cells,
            )
            lastCellValue = cells[len(cells)-1].value
            //fmt.Printf("%+v\n", cells[len(cells)-1])
        }
        direction *= -1
    }
    fmt.Println(lastCellValue)
}
