package tools

import (
    "math/rand"
    "strconv"
)

func Shuffle[T any](slice []T) []T {
    for i := range slice {
        j := rand.Intn(i + 1)
        slice[i], slice[j] = slice[j], slice[i]
    }
    return slice
}

func ToString(value interface{}) string {
    switch value.(type) {
    case int:
        return strconv.Itoa(value.(int))
    default:
        panic("Unknown type")
    }
}

func Repeated(value int, list []int) []int {
    repeated := make([]int, 0)
    for i, v := range list {
        if v == value { 
            repeated = append(repeated, i)
        }
    }
    return repeated
}

func Max(list []int) int {
    max := list[0]
    for _, v := range list {
        if v > max { max = v }
    }
    return max
}

func SortByIndex(list []int) []int {
    indexes := make([]int, len(list))
    for i := 0; i < len(list); i++ {
        indexes[i] = i
    }
    for i := 0; i < len(list); i++ {
        for j := 0; j < len(list)-1-i; j++ {
            if list[indexes[j]] < list[indexes[j+1]] {
                indexes[j+1], indexes[j] = indexes[j], indexes[j+1]
            }
        }
    }
    return indexes
}
