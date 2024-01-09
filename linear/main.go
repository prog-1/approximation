package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Point struct {
	x, y float64
}

// Game is a game instance.
type Game struct {
	width, height int
	points        []Point
}

// NewGame returns a new Game instance.
func NewGame(width, height int, pts []Point) *Game {
	return &Game{
		width:  width,
		height: height,
		points: pts,
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

// Update updates a game state.
func (g *Game) Update() error {
	return nil
}

// Draw renders a game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	for _, p := range g.points {
		vector.DrawFilledCircle(screen, float32(p.x), float32(p.y), 3, color.RGBA{255, 255, 0, 0}, false)
	}
	a, b := approximation(g.points)
	var x float64
	for ; x < float64(g.width); x += 0.1 {
		screen.Set(int(x), int(a*x+b), color.White)
	}
}

func approximation(pts []Point) (a, b float64) {
	var sumX, sumY, sumXY, sumX2 float64
	for _, p := range pts {
		sumX += p.x
		sumY += p.y
		sumXY += p.x * p.y
		sumX2 += p.x * p.x
	}
	n := float64(len(pts))
	a = (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	b = (sumY - a*sumX) / n
	return
}

func main() {
	pts := []Point{{100, 300}, {300, 450}, {500, 150}, {700, 300}}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame(screenWidth, screenHeight, pts)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
