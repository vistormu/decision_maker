package dice

import (
    "math/rand"
)

type Dice struct {
    probabilities []float64
}
func New() Dice {
    return Dice{probabilities: []float64{0.1, 0.2, 0.3, 0.2, 0.1, 0.1}}
}
func (self Dice) Roll() int {
    var sum float64 = 0.0
    var r float64 = rand.Float64()
    for i := 0; i < len(self.probabilities); i++ {
        sum += self.probabilities[i]
        if r < sum {
            return i + 1
        }
    }
    return len(self.probabilities)  
}
