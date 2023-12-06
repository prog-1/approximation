package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type coord struct {
	X, Y float64
}

type Game struct {
	width, height int
	dots          []coord
}

const (
	screenWidth  = 1920
	screenHeight = 1080
)

func NewGame(width, height int, d []coord) *Game {
	return &Game{
		width:  width,
		height: height,
		dots:   d,
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Derivative Visualization")
	dots := []coord{{100, 300}, {200, 100}, {300, 200}, {400, 400}}
	g := NewGame(screenWidth, screenHeight, dots)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

func approximate(dots []coord) (a, b float64) {
	n := float64(len(dots))
	var sumX, sumX2, sumY, sumXY float64
	for _, j := range dots {
		sumX += j.X
		sumX2 += j.X * j.X
		sumY += j.Y
		sumXY += j.X * j.Y
	}
	a = (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	b = (sumY - a*sumX) / n
	return a, b
}

func (g Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, j := range g.dots {
		vector.DrawFilledCircle(screen, float32(j.X), float32(j.Y), 3, color.RGBA{255, 255, 0, 255}, false)
	}
	a, b := approximate(g.dots)
	vector.StrokeLine(screen, 0, float32(b), float32(g.width), float32(a*float64(g.width)+b), 3, color.White, false)
}

func (g *Game) Layout(width, height int) (screenWidth, screenHeight int) {
	return g.width, g.height
}