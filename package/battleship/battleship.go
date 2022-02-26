package battleship

import (
	"math"
	"math/rand"
	"time"
)

type Position struct {
	X int
	Y int
}

var (
	Board [][]rune
	Ship1 Position
	Ship2 Position
)

func NewGame(boardX, boardY int) (ship1 Position, ship2 Position) {

	Board = make([][]rune, boardX)
	for i := range Board {
		Board[i] = make([]rune, boardY)
	}

	rand.Seed(time.Now().UnixNano())
	Ship1 = Position{
		X: rand.Intn(boardX),
		Y: rand.Intn(boardY),
	}

	Ship2 = Position{
		X: rand.Intn(boardX),
		Y: rand.Intn(boardY),
	}

	for Ship2.X == Ship1.X && Ship2.Y == Ship1.Y {
		Ship2 = Position{
			X: rand.Intn(boardX),
			Y: rand.Intn(boardY),
		}
	}

	return Ship1, Ship2
}

func PlayerMove(move Position) string {

	if move.X >= len(Board) || move.Y >= len(Board[0]) {
		return "out"
	}
	diffShip1 := math.Abs(float64((move.X - Ship1.X) + (move.Y - Ship1.Y)))
	diffShip2 := math.Abs(float64((move.X - Ship2.X) + (move.Y - Ship2.Y)))

	if diffShip1 == 0 {
		Ship1 = Position{99, 99}
		Board[move.X][move.Y] = 'H'
		return "hit"
	} else if diffShip2 == 0 {
		Ship2 = Position{99, 99}
		Board[move.X][move.Y] = 'H'
		return "hit"
	}

	Board[move.X][move.Y] = 'M'
	if diffShip1 < diffShip2 {
		if diffShip1 <= 2 {
			return "hot"
		} else if diffShip1 <= 4 {
			return "warm"
		} else {
			return "cold"
		}
	} else {
		if diffShip2 <= 2 {
			return "hot"
		} else if diffShip2 <= 4 {
			return "warm"
		} else {
			return "cold"
		}
	}
}
