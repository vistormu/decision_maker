import time
import os

from dataclasses import dataclass


@dataclass
class UI:
    scores: list[str]
    board: list[str]
    events: list[str]

    def render(self) -> None:
        self.clear()

        for score in self.scores:
            print(score)

        first_line: str = ''.join(self.board[:6]) + " " + self.scores[0]
        second_line: str = self.board[19] + " " * (4*3) + self.board[6] + " " + self.scores[1]

        print(first_line)
        print(second_line)
        print(self.board[18] + " " * (4*3) + self.board[7])
        print(self.board[17] + " " * (4*3) + self.board[8])
        print(self.board[16] + " " * (4*3) + self.board[9])
        print(''.join(reversed(list(self.board[10:16]))))
        print()

        for event in self.events:
            print(event)

    def set_board(self, board: list[str]) -> None:
        self.board = board

    def append_score(self, score: str) -> None:
        self.scores.append(score)

    def append_event(self, event: str) -> None:
        self.events.append(event)

    def clear(self) -> None:
        os.system("clear")

    def reset(self) -> None:
        self.scores.clear()
        self.events.clear()

    def reset_scores(self) -> None:
        self.scores.clear()

    def wait(self, wait_time: float) -> None:
        time.sleep(wait_time)

    def clear_line(self) -> None:
        print('\033[F\033[K', end='')
