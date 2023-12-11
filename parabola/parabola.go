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
	x, y float64
}

func PorabolicApproximation(points []point) func(x float64) float64 {
	var sx, sy, sxx, sxxx, sxxxx, sxy, sxxy float64
	for i := 0; i < len(points); i++ {
		sx += points[i].x
		sy += points[i].y
		sxx += points[i].x * points[i].x
		sxxx += points[i].x * points[i].x * points[i].x
		sxxxx += points[i].x * points[i].x * points[i].x * points[i].x
		sxy += points[i].x * points[i].y
		sxxy += points[i].x * points[i].x * points[i].y
	}
	n := float64(len(points))

	d := sxxxx*(sxx*n-sx*sx) - sxxx*(sxxx*n-sx*sxx) + sxx*(sxx*sx-sxx*sxx)
	da := sxxy*(sxx*n-sx*sx) - sxy*(sxxx*n-sx*sxx) + sy*(sxx*sx-sxx*sxx)
	db := sxxxx*(sxy*n-sy*sx) - sxxx*(sxxy*n-sy*sxx) + sxx*(sxxy*sx-sxy*sxx)
	dc := sxxxx*(sxx*sy-sx*sxy) - sxxx*(sxxx*sy-sx*sxxy) + sxx*(sxx*sxy-sxx*sxxy)

	a := da / d
	b := db / d
	c := dc / d

	return func(x float64) float64 { return a*x*x + b*x + c*0.1 }
}

//####################################### DRAWING ###########################################

func (g *game) Draw(screen *ebiten.Image) {

	image := ebiten.NewImage(screenWidth, screenHeight)

	points := []point{
		{400, 200},
		{500, 300},
		{300, 300},
		{400, 500},
	}

	DrawPoints(image, points)
	DrawParabola(screen, PorabolicApproximation(points))

	// Draw on screen
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(image, op)
}

func DrawPoints(image *ebiten.Image, points []point) {
	for i := 0; i < len(points); i++ {
		vector.DrawFilledCircle(image, float32(points[i].x), float32(points[i].y), 5, color.RGBA{0, 255, 255, 255}, true)
	}
}

func DrawParabola(image *ebiten.Image, f func(x float64) float64) {
	for x := 0; x < screenWidth; x++ {
		y := f(float64(x))
		drawPixel(image, x, int(y), color.White)
	}
}

func drawPixel(img *ebiten.Image, x, y int, clr color.Color) {
	if x >= 0 && x < screenWidth && y >= 0 && y < screenHeight {
		img.Set(x, y, clr)
	}
}

//####################################### MAIN ###########################################

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Parabolic approximation")
	if err := ebiten.RunGame(&game{}); err != nil {
		panic(err)
	}
}

type game struct{}

func (g *game) Update() error {
	return nil
}

func (g *game) Layout(width, weight int) (screenWidth, screenHeight int) {
	return width, weight
}
