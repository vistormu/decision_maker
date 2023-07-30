package tools

import (
    "testing"
    "fmt"
)

func TestRepeated(t *testing.T) {
    list := []int{1, 2, 3, 4, 5, 6, 7, 8, 1, 1}
    repeated := Repeated(1, list)
    if len(repeated) != 3 {
        t.Errorf("Expected 3, got %d", len(repeated))
    }
    if repeated[0] != 0 {
        t.Errorf("Expected 0, got %d", repeated[0])
    }
    if repeated[1] != 8 {
        t.Errorf("Expected 8, got %d", repeated[1])
    }
    if repeated[2] != 9 {
        t.Errorf("Expected 9, got %d", repeated[2])
    }
}

func TestSortByIndex(t *testing.T) {
    list := []int{1, 7, 4, 9, 3, 5, 2, 8, 6, 0}
    sorted := SortByIndex(list)
    fmt.Println(sorted)
    expected := []int{3, 7, 1, 8, 5, 2, 4, 6, 0, 9}
    if len(sorted) != len(expected) {
        t.Errorf("Expected 10, got %d", len(sorted))
    }
    for i, v := range sorted {
       if v != expected[i] {
            t.Errorf("Expected %d, got %d", expected[i], v)
        }
    }
}
