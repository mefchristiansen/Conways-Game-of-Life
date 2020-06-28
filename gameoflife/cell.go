package gameoflife

import "github.com/hajimehoshi/ebiten"

type Cell struct {
	row   int
	col   int
	state bool
}

func (c *Cell) Draw(boardImage *ebiten.Image) {

}
