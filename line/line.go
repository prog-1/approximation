package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type point struct {
	x, y float32
}

func approximate(points []point) func(x float32) float32 {
	var sumx, sumy, sumxy, sumxx float32
	for i := 0; i < len(points); i++ {
		sumx += points[i].x
		sumy += points[i].y
		sumxy += points[i].x * points[i].y
		sumxx += points[i].x * points[i].x
	}
	n := float32(len(points))
	a := (n*sumxy - sumx*sumy) / (n*sumx*sumx - sumxx)
	b := (sumy - a*sumx) / n
	return func(x float32) float32 { return a*x + b }
}

func update(screen *ebiten.Image) error {

	image := ebiten.NewImage(screenWidth, screenHeight)

	points := []point{
		{500, 400},
		{100, 300},
		{600, 500},
		{400, 200},
	}

	DrawApproximation(image, approximate(points))
	DrawPoints(image, points)

	// Draw on screen
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(image, op)

	return nil
}

func DrawApproximation(image *ebiten.Image, f func(x float32) float32) {
	vector.StrokeLine(image, 0, f(0), screenWidth, f(screenWidth), 2, color.RGBA{0, 255, 0, 255}, true)
}

func DrawPoints(image *ebiten.Image, points []point) {
	for i := 0; i < len(points); i++ {
		vector.DrawFilledCircle(image, points[i].x, points[i].y, 5, color.RGBA{0, 255, 255, 255}, true)
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Linear approximation")
	if err := ebiten.RunGame(&game{}); err != nil {
		panic(err)
	}
}

type game struct{}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	if err := update(screen); err != nil {
		panic(err)
	}
}

func (g *game) Layout(width, weight int) (screenWidth, screenHeight int) {
	return width, weight
}
