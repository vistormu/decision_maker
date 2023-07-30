from decision_maker import Game, Player


def main() -> None:
    game = Game([Player("Castillo ambulante"), 
                 Player("Princesa Mononoke"),
                 Player("Juno"),
                 Player("What if"),
                 Player("Brokeback Mountain"),
                 ])

    steps = 10
    for _ in range(steps):
        game.step()

    winner: Player = game.get_winner()
    print(f'El ganador es {winner.name} con {winner.score} puntos')


if __name__ == "__main__":
    main()
