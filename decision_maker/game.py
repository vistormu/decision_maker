import random
from collections import Counter

from .player import Player
from .board import Board
from .dice import Dice
from .event import Event
from .ui import UI

SPEED = 10

class Game:
    def __init__(self, players: list[Player]) -> None:
        if len(players) > 6:
            raise Exception("The game only supports up top 6 players")

        self.players = players
        starting_player: int = random.randint(0, len(players))
        self.players = self.players[starting_player:] + self.players[:starting_player]

        self.board = Board(20)
        self.dice = Dice()
        self.ui: UI = UI([], [], [])

    def step(self) -> None:
        self.trigger_bowser_event()
        for i, player in enumerate(self.players):
            self.ui.reset()
            self.update_board()
            self.ui.append_event(f'TURNO DE {player.name.upper()}')
            self.ui.render()
            self.ui.wait(1.0 / SPEED)

            dice_roll = self.dice.roll()

            self.ui.append_event(f'{player.name} ha sacado un {dice_roll}')
            self.ui.render()
            self.ui.wait(1.0 / SPEED)

            self.move(player, dice_roll)

            if random.random() > 0.8:
                self.ui.append_event(f'{player.name} se ha librado de un evento')
                self.ui.render()
                self.ui.wait(3.0 / SPEED)
                continue

            event: Event = self.board.get_event(player.position)
            self.ui.append_event(f'{player.name} {event.description}')
            self.ui.render()

            if event.score > 0:
                self.ui.append_event(f'{player.name} ha ganado {event.score} puntos')
                self.ui.render()
            elif event.score < 0:
                self.ui.append_event(f'{player.name} ha perdido {abs(event.score)} puntos')
                self.ui.render()

            if event.absolute_position != -1:
                player.position = event.absolute_position
                self.update_board()
                self.ui.render()
            elif event.relative_position != 0:
                self.move(player, event.relative_position)

            player.score += event.score

            self.ui.scores[i] = f'{player.name}: {player.score}'
            self.ui.render()
            self.ui.wait(3.0 / SPEED)

    def move(self, player: Player, positions: int) -> None:
        if positions + player.position >= self.board.size:
            player.score += 1
            self.ui.append_event(f'{player.name} ha dado la vuelta al tablero y ha ganado 1 punto')

        for _ in range(abs(positions)):
            if positions < 0:
                player.position = (player.position - 1) % self.board.size
            else:
                player.position = (player.position + 1) % self.board.size
            self.update_board()
            self.ui.render()
            self.ui.wait(0.3 / SPEED)

    def update_board(self) -> None:
        board_ui: list[str] = self.board.get_ui()
        colors: list[str] = ['\033[31m', '\033[32m', '\033[34m', '\033[35m', '\033[36m', '\033[37m']
        self.ui.reset_scores()
        for player, color in zip(self.players, colors):
            player_ui: str = f'{color}[{player.name[0]}]\033[0m'
            board_ui[player.position] = player_ui
            self.ui.append_score(f'{player.name}: {player.score}')

        counter: Counter = Counter([player.position for player in self.players])
        duplicates: list[int] = [position for position, count in counter.items() if count > 1]
        for duplicate in duplicates:
            board_ui[duplicate] = '\033[93m[o]\033[0m'

        self.ui.set_board(board_ui)

    def get_winner(self) -> Player:
        scores: list[int] = [player.score for player in self.players]
        if len(set(scores)) == 1:
            return random.choice(self.players)

        return max(self.players, key=lambda player: player.score)

    def trigger_bowser_event(self) -> None:
        if random.random() < 0.5:
            return

        sorted_players: list[Player] = sorted(self.players, key=lambda player: player.score)
        if sorted_players[-1].score - sorted_players[-2].score <= 2:
            return

        self.ui.append_event('\nBOWSER HA APARECIDO')
        self.ui.render()
        self.ui.wait(1.0 / SPEED)

        random_number: int = random.randint(0, 2)
        self.ui.append_event(f'BOWSER HA ROBADO {random_number} PUNTOS A {sorted_players[-1].name.upper()}')
        self.ui.render()
        self.ui.wait(3.0 / SPEED)

        sorted_players[-1].score -= random_number
