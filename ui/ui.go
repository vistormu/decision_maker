package ui

import (
    "fmt"
    "time"
    "strings"
    "decision_maker/tools"
    "decision_maker/player"
    "decision_maker/board"
)

type StateType int

const (
    _ = iota
    NEW_ROUND
    NEW_TURN
    ROLL_DICE
    MOVE_PLAYER
    EVENT_FREED
    STARTING_BOX_CROSSED
    EVENT
    END_TURN
    WINNER
    BOWSER_EVENT
)

const END = "\033[0m"

var colors = []string{
    "\033[31m",
    "\033[32m",
    "\033[33m",
    "\033[34m",
    "\033[35m",
    "\033[36m",
    "\033[93m",
}


type UI struct {
    speed float64
    board board.Board
    players []*player.Player
    lines []string
}
func New(speed float64, board board.Board, players []*player.Player) UI {
    return UI{speed, board, players, []string{}}
}

// Public
func (self *UI) SetState(state StateType, args... interface{}) {
    switch state {
    case NEW_ROUND:
        self.reset()
        self.addLine("Round " + tools.ToString(args[0].(int)) + "/" + tools.ToString(args[1].(int)))
        self.draw()
    case NEW_TURN:
        self.addLine("")
        self.addLine(self.players[args[0].(int)].Name + "'s turn")
        self.wait(1)

    case ROLL_DICE:
        self.addLine("- " + self.players[args[0].(int)].Name + " rolled a " + tools.ToString(args[1].(int)))
        self.draw()
        self.wait(1)

    case STARTING_BOX_CROSSED:
        self.addLine("- " + self.players[args[0].(int)].Name + " receives a point for reaching the start of the board")
        self.draw()

    case MOVE_PLAYER:
        self.draw()
        self.wait(0.3)

    case EVENT_FREED:
        self.addLine("- " + self.players[args[0].(int)].Name + " has freed of an event")
        self.draw()
        self.wait(2)

    case EVENT:
        player := self.players[args[0].(int)].Name
        event := args[1].(board.Event)
        self.addLine("- " + player + " " + event.Description)
        self.addLine("- " + player + " receives " + tools.ToString(event.Score) + " points")
        self.draw()
        self.wait(1)

    case END_TURN:
        self.draw()
        self.wait(2)

    case WINNER:
        self.reset()
        self.addLine(args[0].(string) + " wins the game!")
        self.draw()

    case BOWSER_EVENT:
        if len(args) == 0 {
            self.addLine("BOWSER HAS APPEARED!")
            self.draw()
            self.wait(1)
        } else {
            self.addLine("- BOWSER HAS STOLEN " + tools.ToString(args[0].(int)) + " POINTS FROM " + self.players[args[1].(int)].Name)
            self.draw()
            self.wait(2)
        }
    }
}

// Private
func (self UI) wait(seconds float64) {
    time.Sleep(time.Duration(seconds / self.speed * 1000) * time.Millisecond)
}
func (self *UI) addLine(line string) {
    self.lines = append(self.lines, line)
}
func (self *UI) reset() {
    self.lines = []string{}
}
func (self UI) draw() {
    clear()
    self.drawLines()
    self.drawBoard()
}
func clear() {
    fmt.Print("\033[2J")
}
func (self UI) drawLines() {
    for _, line := range self.lines {
        fmt.Println(line)
    }
    fmt.Println()
}
func (self UI) drawBoard() {
    boardLine := make([]string, self.board.Size)
    for i := 0; i < self.board.Size; i++ {
        boardLine[i] = self.getCell(i)
    }
    boardUI := make([]string, 6)
    boardUI[0] = strings.Join(boardLine[0:6], "")
    boardUI[1] = boardLine[19] + "            " + boardLine[6]
    boardUI[2] = boardLine[18] + "            " + boardLine[7]
    boardUI[3] = boardLine[17] + "            " + boardLine[8]
    boardUI[4] = boardLine[16] + "            " + boardLine[9]
    lastLine := ""
    for i := 15; i > 9; i-- {
        lastLine += boardLine[i]
    }
    boardUI[5] = lastLine

    scores := make([]int, len(self.players))
    for i, player := range self.players {
        scores[i] = player.Score
    }

    scoresByIndex := tools.SortByIndex(scores)
    for i := 0 ; i < len(scoresByIndex); i++ {
        player := self.players[scoresByIndex[i]]
        boardUI[i] += " " + player.Name + ": " + tools.ToString(player.Score)
    }

    for _, line := range boardUI {
        fmt.Println(line)
    }
}
func (self UI) getCell(position int) string {
    playerInCell := false
    playerString := ""
    for i, player := range self.players {
        if player.Position == position {
            if playerInCell {
                playerString = fmt.Sprintf("[%s*%s]", colors[6], END)
                break
            }
            playerInCell = true
            playerString += fmt.Sprintf("[%s%s%s]", colors[i], player.Name[0:1], END)
        }
    }
    if playerInCell {
        return playerString
    }

    if position == 0 {
        return "[S]"
    }
    return "[ ]"
}
