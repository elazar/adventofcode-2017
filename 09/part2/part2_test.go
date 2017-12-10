package main

import (
    "./part2"
    "testing"
)

func TestGetScore(t *testing.T) {
    cases := map[string]int{
        "<>": 0,
        "<random characters>": 17,
        "<<<<>": 3,
        "<{!>}>": 2,
        "<!!>": 0,
        "<!!!>>": 0,
        "<{o\"i!a,<{i<a>": 10,
    }
    for input, expected := range cases {
        actual := part2.GetScore([]byte(input))
        if actual != expected {
            t.Errorf("%s: %d vs %d\n", input, expected, actual)
        }
    }
}
