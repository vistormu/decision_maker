# Decision Maker
Have you ever felt like you don't want to decide between a series of options? Decision Maker helps your indecision by running a board game!

Decision maker is writen in Python and Go, being Go the improved version.

```
package main

import (
    "decision_maker/game"
)

func main() {
    players := []string{"Alice", "Bob", "Charlie", "Dave", "Eve"}
    speed := 1.0
    rounds := 10

    game := game.New(players, speed)
    game.Play(rounds)
}
```
