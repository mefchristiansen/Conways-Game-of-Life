package gameoflife

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

// The game board
type Board struct {
	// The board dimensions (i.e. num rows and num columns)
	dimension int
	// The scale of each cell relative to the game screen
	scale int
	// The 2D array of cells
	cells [][]Cell
}

// NewBoard generates a new Board with giving a size.
func NewBoard(dimension int, scale int, initialStateProbability float32) (*Board, error) {
	b := &Board{
		dimension: int(dimension),
		scale:     int(scale),
		cells:     initBoard(dimension, initialStateProbability),
	}

	return b, nil
}

// Initializes the game board and randomizes the cell state of the initial generation
func initBoard(dimension int, initialStateProbability float32) [][]Cell {
	cells := make([][]Cell, dimension)

	// Seed the random number generator to get a different initial state every time
	rand.Seed(time.Now().UTC().UnixNano())

	for row := 0; row < dimension; row++ {
		cells[row] = make([]Cell, dimension)

		for col := 0; col < dimension; col++ {
			cells[row][col] = Cell{
				row:   row,
				col:   col,
				state: rand.Float32() < initialStateProbability,
			}
		}
	}

	return cells
}

// Update updates the board state.
func (b *Board) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		b.flip(x, y)
	}

	b.generateNextGeneration()

	return nil
}

// Sets the state of the clicked cell and its neighboring cells to alive
func (b *Board) flip(x int, y int) {
	// Scale the mouse click x and y pixel coordinates to the board dimensions
	row := int(x / b.scale)
	col := int(y / b.scale)

	// Iterate through the current cell's neighbors
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			diffRow := row + dr
			diffCol := col + dc

			if diffRow >= 0 && diffCol >= 0 && diffRow < b.dimension && diffCol < b.dimension {
				// If the cell is a valid cell on the board, set it to being alive
				b.cells[diffRow][diffCol].state = true
			}
		}
	}
}

// Generate the next iteration of cells on the board using the current generation
func (b *Board) generateNextGeneration() {
	// Generate an initial empty state
	nextGeneration := *b.emptyGeneration()

	// Iterate through all cells on the board
	for row := 0; row < b.dimension; row++ {
		for col := 0; col < b.dimension; col++ {
			cell := &b.cells[row][col]
			numNeighbors := b.countCellNeighbors(cell)
			alive := cell.state

			if alive && (numNeighbors == 2 || numNeighbors == 3) {
				// If a cell is alive and has 2 or 3 neighbors, it survives
				nextGeneration[row][col].state = true
			} else if !alive && numNeighbors == 3 {
				// If a cell is dead and has 3 neighbors, it becomes alive
				nextGeneration[row][col].state = true
			} else {
				// All other live cells die in the next generation. Similarly, all other dead cells stay dead.
				nextGeneration[row][col].state = false
			}
		}
	}

	b.cells = nextGeneration
}

// Generates an empty board state used to determine the next generation of cells
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

// Counts the number of alive neighboring cells of the input cell
func (b *Board) countCellNeighbors(c *Cell) int {
	neighbors := 0

	// Iterate through the current cell's neighbors
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				// Ignore the input cell's own state
				continue
			}

			// Increment the number of valid alive neighbors
			neighbors += b.validNeighbor(c.row+dr, c.col+dc)
		}
	}

	return neighbors
}

// Returns either a 1 or 0 if the neighbor is either valid and alive or not
func (b *Board) validNeighbor(row, col int) int {
	if !(row >= 0 && col >= 0 && row < b.dimension && col < b.dimension) {
		// If the cell is out of bounds of the board, this is an invalid neighbor
		return 0
	}

	if b.cells[row][col].state {
		// If the cell is alive, then this is a valid alive neighbor
		return 1
	} else {
		// If the cell is dead, this is not a valid alive neighbor
		return 0
	}
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(boardImage *ebiten.Image) {
	for row := 0; row < b.dimension; row++ {
		for col := 0; col < b.dimension; col++ {
			c := b.cells[row][col]
			c.Draw(boardImage, b.scale)
		}
	}
}
