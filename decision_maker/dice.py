import random
from typing import NamedTuple


class Dice(NamedTuple):
    probabilities: list[float] = [0.1, 0.2, 0.3, 0.2, 0.1, 0.1]

    def roll(self) -> int:
        return random.choices(range(1, 7), weights=self.probabilities)[0]
