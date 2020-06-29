package gameoflife

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Cell struct {
	row   int
	col   int
	state bool
}

func (c *Cell) Draw(boardImage *ebiten.Image, scale int) {
	var color color.Gray16

	if c.state {
		color = aliveColor
	} else {
		color = deadColor
	}

	ebitenutil.DrawRect(boardImage,
		float64(c.row*scale),
		float64(c.col*scale),
		float64(scale),
		float64(scale),
		color,
	)
}
