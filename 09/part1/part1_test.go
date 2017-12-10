package main

import (
    "./part1"
    "testing"
)

func TestGetScore(t *testing.T) {
    cases := map[string]int{
        "{}": 1,
        "{{{}}}": 6,
        "{{}{}}": 5,
        "{{{},{},{{}}}}": 16,
        "{<a>,<a>,<a>,<a>}": 1,
        "{{<a>},{<a>},{<a>},{<a>}}": 9,
        "{{<!!>},{<!!>},{<!!>},{<!!>}}": 9,
        "{{<a!>},{<a!>},{<a!>},{<ab>}}": 3,
    }
    for input, expected := range cases {
        actual := part1.GetScore([]byte(input))
        if actual != expected {
            t.Errorf("%s: %d vs %d\n", input, expected, actual)
        }
    }
}
