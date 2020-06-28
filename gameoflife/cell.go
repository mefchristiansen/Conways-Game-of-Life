package gameoflife

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Cell struct {
	row   int
	col   int
	state bool
}

func (c *Cell) Draw(boardImage *ebiten.Image, scale int) {
	ebitenutil.DrawRect(boardImage,
		float64(c.row*scale),
		float64(c.col*scale),
		float64(scale),
		float64(scale),
		aliveColor,
	)
}
