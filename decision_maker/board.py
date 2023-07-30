from .event import Event, Events


class Board:
    def __init__(self, size: int) -> None:
        self.size = size
        self._events: list[Event] = Events().get(size)

    def get_event(self, position: int) -> Event:
        return self._events[position]

    def get_ui(self) -> list[str]:
        board_ui: list[str] = ["[ ]" for _ in range(self.size)]
        board_ui[0] = "[X]"
        return board_ui
