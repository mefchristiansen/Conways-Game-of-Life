package gameoflife

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	ScreenDimension int = 600
	scale           int = 4
)

type Game struct {
	board      *Board
	boardImage *ebiten.Image
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	g := &Game{}
	var err error
	g.board, err = NewBoard(ScreenDimension, scale)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenDimension, ScreenDimension
}

func (g *Game) Update(*ebiten.Image) error {
	if err := g.board.Update(); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		w, h := ScreenDimension, ScreenDimension
		g.boardImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	}
	screen.Fill(deadColor)
	g.board.Draw(g.boardImage)
	screen.DrawImage(g.boardImage, &ebiten.DrawImageOptions{})
}

func main() {
	return
}
