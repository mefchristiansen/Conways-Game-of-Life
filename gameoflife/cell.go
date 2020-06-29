package gameoflife

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Cell represents a cell on the board in the game of life
type Cell struct {
	row int
	col int
	// State indicating whether the cell is alive or dead
	state bool
}

// Draws a cell given its current state
func (c *Cell) Draw(boardImage *ebiten.Image, scale int) {
	var color color.Gray16

	if c.state {
		// If alive, the cell fill is white
		color = aliveColor
	} else {
		// Otherwise, it is black (i.e. dead)
		color = deadColor
	}

	// Draws a recentangle (representing a cell) to the game screen
	ebitenutil.DrawRect(boardImage,
		float64(c.row*scale),
		float64(c.col*scale),
		float64(scale),
		float64(scale),
		color,
	)
}
