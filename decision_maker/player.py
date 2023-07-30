from dataclasses import dataclass


@dataclass
class Player:
    name: str
    score: int = 0
    position: int = 0

    def __repr__(self):
        return f"{self.name}: {self.score}"
