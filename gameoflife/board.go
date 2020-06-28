package gameoflife

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

type Board struct {
	dimension int
	scale     int
	cells     [][]Cell
}

// NewBoard generates a new Board with giving a size.
func NewBoard(dimension int, scale int) (*Board, error) {
	b := &Board{
		dimension: int(dimension),
		scale:     int(scale),
		cells:     initBoard(dimension),
	}

	return b, nil
}

func (b *Board) Update() error {
	return nil
}

func (b *Board) Draw(boardImage *ebiten.Image) {
	for row := 0; row < b.dimension; row++ {
		for col := 0; col < b.dimension; col++ {

			c := b.cells[row][col]
			if c.state {
				c.Draw(boardImage, b.scale)
			}
		}
	}
}

func initBoard(dimension int) [][]Cell {
	cells := make([][]Cell, dimension)

	rand.Seed(time.Now().UTC().UnixNano())

	for row := 0; row < dimension; row++ {
		cells[row] = make([]Cell, dimension)

		for col := 0; col < dimension; col++ {
			cells[row][col] = Cell{
				row:   row,
				col:   col,
				state: rand.Float32() < 0.5,
			}
		}
	}

	return cells
}
