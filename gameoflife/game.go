package gameoflife

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	// The game screen size (in pixels)
	ScreenDimension int = 600
	// The number of rows and columns (num rows == num columns)
	boardDimensions int = 100
	// The probability that a cell is alive in the initial generation
	initialStateProbability float32 = 0.5
)

// Game represents a game state.
type Game struct {
	board      *Board
	boardImage *ebiten.Image
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	g := &Game{}
	var err error
	g.board, err = NewBoard(
		boardDimensions,
		ScreenDimension/boardDimensions,
		initialStateProbability,
	)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenDimension, ScreenDimension
}

// Updates the current game state.
func (g *Game) Update(*ebiten.Image) error {
	if err := g.board.Update(); err != nil {
		return err
	}
	return nil
}

// Draws the current game state to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		w, h := ScreenDimension, ScreenDimension
		g.boardImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	}
	screen.Fill(deadColor)
	g.board.Draw(g.boardImage)
	screen.DrawImage(g.boardImage, &ebiten.DrawImageOptions{})
}
