package main

import (
	"log"

	"github.com/crowllx/geometry"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

type Game struct {
	shapes []geometry.Shape
	player geometry.Shape
}

func NewGame() *Game {
	return &Game{
		[]geometry.Shape{
			geometry.NewCircle(geometry.NewVector(400, 200), 20),
			geometry.NewCircle(geometry.NewVector(100, 200), 20),
			geometry.NewRect(200, 250, 250, 300),
		},
		geometry.NewRect(0, 0, 50, 50),
	}
}

func DrawShape(s geometry.Shape, screen *ebiten.Image) {
	switch s.(type) {
	case geometry.Circle:
		circle := s.(geometry.Circle)
		center := circle.Center()
		vector.StrokeCircle(
			screen,
			float32(center.X),
			float32(center.Y),
			float32(circle.Radius()),
			2,
			colornames.Green,
			false,
		)
		bb := circle.BB()
		vector.StrokeRect(screen, float32(bb.L), float32(bb.T), float32(bb.R-bb.L), float32(bb.B-bb.T), 2, colornames.Darkcyan, false)
	case geometry.Rect:
		rect := s.(geometry.Rect)
		dx, dy := rect.Bounds()
		vector.StrokeRect(
			screen,
			float32(rect.Min.X),
			float32(rect.Min.Y),
			float32(dx),
			float32(dy),
			2,
			colornames.Azure,
			false,
		)

	}

}

// Draw implements ebiten.Game.
func (g *Game) Draw(screen *ebiten.Image) {
	for _, shape := range g.shapes {
		DrawShape(shape, screen)
	}
	DrawShape(g.player, screen)
	for _, shape := range g.shapes {
		if g.player.BB().Contains(shape.BB()) {
			ebitenutil.DebugPrint(screen, "intersection")
		}
	}
}

// Layout implements ebiten.Game.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 1600, 900
}

// Update implements ebiten.Game.
func (g *Game) Update() error {
	x, y := 0.0, 0.0
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		x -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		x += 1
	}

	velocity := geometry.NewVector(x, y).Mult(2)
	g.player = g.player.(geometry.Rect).Scale(velocity)

	return nil
}

var _ ebiten.Game = &Game{}

func main() {
	g := NewGame()

	ebiten.SetWindowSize(1600, 900)
	ebiten.SetWindowTitle("shapes")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
