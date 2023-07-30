package dice

import (
    "testing"
    "math"
)

func TestRoll(t *testing.T) {
    dice := New()

    rolls := make([]int, 6)
    times := 1000
    for i := 0; i < times; i++ {
        index := dice.Roll() - 1
        rolls[index]++
    }

    probabilities := make([]float64, 6)
    for i := 0; i < len(probabilities); i++ {
        probabilities[i] = math.Round(float64(rolls[i])/float64(times)*10)/10
    }

    tests := []struct {
        value float64
        expected float64
    }{
        {probabilities[0], 0.1},
        {probabilities[1], 0.2},
        {probabilities[2], 0.3},
        {probabilities[3], 0.2},
        {probabilities[4], 0.1},
        {probabilities[5], 0.1},
    }

    for _, test := range tests {
        if test.value != test.expected {
            t.Errorf("Expected %f, got %f", test.expected, test.value)
        }
    }
}
