package game

import (
    "math"
    "math/rand"
    "decision_maker/player"
    "decision_maker/board"
    "decision_maker/dice"
    "decision_maker/ui"
    "decision_maker/tools"
)

type Game struct {
    Players []*player.Player
    Board board.Board
    Dice dice.Dice
    UI ui.UI
}
func New(players []string, speed float64) Game {
    if len(players) < 2 { panic("Not enough players") }
    if len(players) > 6 { panic("Too many players") }

    board := board.New(20)
    dice := dice.New()
    playerList := make([]*player.Player, len(players))
    for i, playerName := range players {
        playerList[i] = player.New(playerName, 0, 0)
    }
    players = tools.Shuffle(players)
    ui := ui.New(speed, board, playerList)
    return Game{playerList, board, dice, ui}
}

func (self Game) Play(rounds int) {
    for i := 1; i <= rounds; i++ {
        self.UI.SetState(ui.NEW_ROUND, i, rounds)
        self.Step()
    }

    winner := self.getWinner()
    self.UI.SetState(ui.WINNER, winner.Name)
}
func (self Game) Step() {
    self.bowserEvent()
    for i, player := range self.Players {
        self.UI.SetState(ui.NEW_TURN, i)

        diceRoll := self.Dice.Roll()
        self.UI.SetState(ui.ROLL_DICE, i, diceRoll)

        if player.Position + diceRoll >= self.Board.Size {
            player.Score++
            self.UI.SetState(ui.STARTING_BOX_CROSSED, i)
        }

        self.movePlayer(player, diceRoll)

        if rand.Float32() > 0.8 {
            self.UI.SetState(ui.EVENT_FREED, i)
            continue
        }
        event := self.Board.Events[player.Position]
        self.UI.SetState(ui.EVENT, i, event)

        if event.AbsolutePosition != -1 {
            player.Position = event.AbsolutePosition
        } else if event.RelativePosition != -1 {
            self.movePlayer(player, event.RelativePosition)
        }

        player.Score += event.Score
        self.UI.SetState(ui.END_TURN)
    }
}
func (self Game) movePlayer(player *player.Player, position int) {
    absPosition := int(math.Abs(float64(position)))
    for i := 0; i < absPosition; i++ {
        if position < 0 { 
            player.Position = (player.Position - 1) % self.Board.Size
        } else { 
            player.Position = (player.Position + 1) % self.Board.Size
        }
        self.UI.SetState(ui.MOVE_PLAYER)
    }
}
func (self Game) bowserEvent() {
    if rand.Float32() > 0.75 { return }
    scores := make([]int, len(self.Players))
    for i, player := range self.Players {
        scores[i] = player.Score
    }
    scoreIndexes := tools.SortByIndex(scores)

    if self.Players[scoreIndexes[0]].Score - self.Players[scoreIndexes[1]].Score<= 2 {
        return
    }

    self.UI.SetState(ui.BOWSER_EVENT)
    randomScore := rand.Intn(2) + 1
    self.Players[scoreIndexes[0]].Score -= randomScore
    self.UI.SetState(ui.BOWSER_EVENT, randomScore, scoreIndexes[0])
}
func (self Game) getWinner() *player.Player {
    scores := make([]int, len(self.Players))
    for i, player := range self.Players {
        scores[i] = player.Score
    }
    maxScore := tools.Max(scores)
    repeatedMaxScores := tools.Repeated(maxScore, scores)

    if len(repeatedMaxScores) == 1 {
        return self.Players[repeatedMaxScores[0]]
    } 

    repeatedShuffledMaxScores := tools.Shuffle(repeatedMaxScores)
    return self.Players[repeatedShuffledMaxScores[0]]
}
