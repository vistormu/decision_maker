package player

type Player struct {
    Name string
    Score int
    Position int
}
func New(name string, score int, position int) *Player {
    return &Player{name, score, position}
}
