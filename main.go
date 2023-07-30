package main

import (
    "decision_maker/game"
)

func main() {
    players := []string{"Alice", "Bob", "Charlie", "Dave", "Eve", "Frank"}
    speed := 1.0
    game := game.New(players, speed)
    rounds := 10
    game.Play(rounds)
}
