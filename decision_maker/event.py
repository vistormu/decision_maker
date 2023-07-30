import random
from typing import NamedTuple


class Event(NamedTuple):
    absolute_position: int
    relative_position: int
    score: int
    description: str


class Events:
    def __init__(self) -> None:
        self._events = [
            Event(-1, 0, 3, "le llega la herencia de su tio que invirtio en criptos"),
            Event(-1, 0, 2, "se cae por las escaleras pero era un sueño y se despierta con una sonrisa"),
            Event(-1, 0, 2, "pide Casa Miguel y le dan una cerveza gratis"),
            Event(-1, 0, 1, "se urgo la nariz y encontro un trozo de chocolate"),
            Event(-1, 0, 1, "se encuentra un billete de 20 euros en el suelo"),
            Event(-1, 0, 1, "se despierta por la noche pensando que tenia que entregar el worbu y se da cuenta de que es sabado"),
            Event(-1, 0, 1, "sale de fiesta y se encuentra a su ex con un tio feo"),
            Event(-1, 0, 1, "se da cuenta de que la oposicion es una mierda y se pone a estudiar"),
            Event(0, 0, 0, "vuelve a la casilla de salida"),
            Event(-1, 2, 0, "se ha quedado sin papel higienico y tiene que ir a comprar y avanzar dos casillas"),
            Event(-1, -2, 0, "se encuentra con el amigo que le habla muy de cerca y tiene que retroceder dos casillas"),
            Event(-1, 3, 0, "cree haber visto a su ex, se pone nervioso y avanza tres casillas"),
            Event(-1, 0, -1, "coge la llamada de un antiguo amigo que le ofrece participar en una estafa piramidal"),
            Event(-1, 0, -1, "se encuentra un billete de 20 euros en el suelo, pero al cogerlo se da cuenta de que es falso"),
            Event(-1, 0, -1, "iba a coger un taxi pero se le adelanta un señor mayor"),
            Event(-1, 0, -1, "se le cae el movil al suelo y se le rompe la pantalla"),
            Event(-1, 0, -1, "iba a coger un hielo para el refresco y se le cae la puerta del congelador en el pie"),
            Event(-1, 0, -2, "estaba en el metro y un mono capuchino de pecho amarillo le robo la cartera"),
            Event(-1, 0, -2, "se pensaba que el clitoris era una parte de la garganta"),
            Event(-1, -4, -3, "grita 'maricon' en la calle y tiene que salir huyendo"),
        ]

    def get(self, n: int) -> list[Event]:
        if n > len(self._events):
            raise ValueError("n cannot be greater than the number of events")

        random.shuffle(self._events)

        return self._events[:n]
