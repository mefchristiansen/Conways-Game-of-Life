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

func (b *Board) Update() error {
	nextGeneration := *b.emptyGeneration()

	for row := 0; row < b.dimension; row++ {
		for col := 0; col < b.dimension; col++ {
			cell := &b.cells[row][col]
			numNeighbors := b.countCellNeighbors(cell)
			alive := cell.state

			if alive && (numNeighbors == 2 || numNeighbors == 3) {
				nextGeneration[row][col].state = true
			} else if !alive && numNeighbors == 3 {
				nextGeneration[row][col].state = true
			} else {
				nextGeneration[row][col].state = false
			}
		}
	}

	b.cells = nextGeneration

	return nil
}

func (b *Board) emptyGeneration() *[][]Cell {
	cells := make([][]Cell, b.dimension)

	for row := 0; row < b.dimension; row++ {
		cells[row] = make([]Cell, b.dimension)

		for col := 0; col < b.dimension; col++ {
			cells[row][col] = Cell{
				row:   row,
				col:   col,
				state: false,
			}
		}
	}

	return &cells
}

func (b *Board) countCellNeighbors(c *Cell) int {
	neighbors := 0

	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}

			neighbors += b.validNeighbor(c.row+dr, c.col+dc)
		}
	}

	return neighbors
}

func (b *Board) validNeighbor(row, col int) int {
	if !(row >= 0 && col >= 0 && row < b.dimension && col < b.dimension) {
		return 0
	}

	if b.cells[row][col].state {
		return 1
	} else {
		return 0
	}
}

func (b *Board) Draw(boardImage *ebiten.Image) {
	for row := 0; row < b.dimension; row++ {
		for col := 0; col < b.dimension; col++ {
			c := b.cells[row][col]
			c.Draw(boardImage, b.scale)
		}
	}
}

// if alive && numNeighbors < 2 {
// 	nextGeneration[row][col].state = false
// }
// if alive && (numNeighbors == 2 || numNeighbors == 3) {
// 	nextGeneration[row][col].state = true
// }
// if alive && numNeighbors > 3 {
// 	nextGeneration[row][col].state = false
// }
// if !alive && numNeighbors == 3 {
// 	nextGeneration[row][col].state = true
// }
