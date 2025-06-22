package main

import (
	"fmt"
	"log"

	"github.com/crowllx/geometry"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

type Game struct {
	shapes      []geometry.Shape
	player      geometry.Shape
	player2     geometry.Shape
	player3     geometry.Shape
	activeShape geometry.Shape
}

func NewGame() *Game {
	return &Game{
		[]geometry.Shape{
			geometry.NewCircle(geometry.NewVector(800, 200), 100),
			geometry.NewCircle(geometry.NewVector(300, 200), 100),
			geometry.NewRect(500, 250, 750, 500),
		},
		geometry.NewCircle(geometry.NewVector(0, 0), 30),
		geometry.NewRect(600, 600, 650, 650),
		geometry.NewConvexPolygon(
			geometry.NewVector(100, 0),
			geometry.NewVector(200, 100),
			geometry.NewVector(0, 100),
		),
		nil,
	}
}

func DrawShape(s geometry.Shape, screen *ebiten.Image) {
	switch s.(type) {
	case *geometry.Circle:
		circle := s.(*geometry.Circle)
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
	case *geometry.Rect:
		rect := s.(*geometry.Rect)
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

	case *geometry.ConvexPolygon:
		cp := s.(*geometry.ConvexPolygon)
		bb := cp.BB()
		vector.StrokeRect(
			screen,
			float32(bb.L),
			float32(bb.T),
			float32(bb.R-bb.L),
			float32(bb.B-bb.T),
			2,
			colornames.Darkcyan,
			false,
		)

		firstVec := cp.Vertices[0]
		prevVec := firstVec
		for _, vec := range cp.Vertices[1:] {
			vector.StrokeLine(
				screen,
				float32(prevVec.X),
				float32(prevVec.Y),
				float32(vec.X),
				float32(vec.Y),
				1,
				colornames.Purple,
				false,
			)
		}
		vector.StrokeLine(
			screen,
			float32(prevVec.X),
			float32(prevVec.Y),
			float32(firstVec.X),
			float32(firstVec.Y),
			1,
			colornames.Purple,
			false,
		)

		vector.DrawFilledCircle(
			screen, float32(cp.Centroid().X), float32(cp.Centroid().Y), 2, colornames.Darkblue, false,
		)

	}

}

// Draw implements ebiten.Game.
func (g *Game) Draw(screen *ebiten.Image) {
	for _, shape := range g.shapes {
		DrawShape(shape, screen)
	}
	DrawShape(g.player, screen)
	DrawShape(g.player2, screen)
	DrawShape(g.player3, screen)
	for _, shape := range g.shapes {
		if g.player.Collides(shape) {
			ebitenutil.DebugPrint(screen, "intersection!")
		}
		if g.player2.Collides(shape) {
			ebitenutil.DebugPrint(screen, "intersection!")
		}
		if g.player3.Collides(shape) {
			ebitenutil.DebugPrint(screen, "intersection!")
		}
	}
}

// Layout implements ebiten.Game.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 1600, 900
}

// Update implements ebiten.Game.
func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.Key1) {
		g.activeShape = g.player
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		g.activeShape = g.player2
	}
	if ebiten.IsKeyPressed(ebiten.Key3) {
		g.activeShape = g.player3
	}

	var x, y float32
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
	g.activeShape.Translate(velocity)

	return nil
}

var _ ebiten.Game = &Game{}

func main() {
	g := NewGame()
	g.activeShape = g.player
	w, h := g.player2.BB().Bounds()
	w2, h2 := g.player2.(*geometry.Rect).Bounds()
	fmt.Printf("%t %t", w == w2, h == h2)

	ebiten.SetWindowSize(1600, 900)
	ebiten.SetWindowTitle("shapes")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
