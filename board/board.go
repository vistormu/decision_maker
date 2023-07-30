package board

import (
    "decision_maker/tools"
)

type Event struct {
    Description string
    Score int
    AbsolutePosition int
    RelativePosition int
}

var events = []Event{
    {"fell from a tree", -3, -1, -1},
    {"fell from the sky", -2, -1, -1},
    {"fell from the roof", -2, -1, -1},
    {"fell from a cliff", -1, -1, -1},
    {"fell from the sofa", -1, -1, -1},
    {"fell from the canape", -1, -1, -1},
    {"was chilling", 0, -1, -1},
    {"was sleeping", 0, -1, -1},
    {"was sleeping", 0, -1, -1},
    {"was sleeping", 0, -1, -1},
    {"was sleeping", 0, -1, -1},
    {"was sleeping", 0, -1, -1},
    {"was sleeping", 0, -1, -1},
    {"was studying", 0, -1, -1},
    {"found a friend", 1, -1, -1},
    {"found a snorlax", 1, -1, -1},
    {"found a quaismodo", 1, -1, -1},
    {"found a mascot", 2, -1, -1},
    {"found a true lover", 2, -1, -1},
    {"found a treasure", 3, -1, -1},
}

type Board struct {
    Size int
    Events []Event
}
func New(size int) Board {
    return Board{size, tools.Shuffle(events)}
}
