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
	a, b, c := approximation(g.points)
	var x float64
	for ; x < float64(g.width); x += 0.1 {
		screen.Set(int(x), int(a*x*x+b*x+c), color.White)
	}
}

func det(a11, a12, a13, a21, a22, a23, a31, a32, a33 float64) float64 {
	return a11*a22*a33 - a11*a23*a32 - a12*a21*a33 + a12*a23*a31 + a13*a21*a32 - a13*a22*a31
}

func approximation(pts []Point) (a, b, c float64) {
	var sumX, sumY, sumXY, sumX2, sumX3, sumX4, sumX2Y float64
	for _, p := range pts {
		sumX += p.x
		sumY += p.y
		sumXY += p.x * p.y
		sumX2 += p.x * p.x
		sumX3 += p.x * p.x * p.x
		sumX4 += p.x * p.x * p.x * p.x
		sumX2Y += p.x * p.x * p.y
	}
	n := float64(len(pts))

	a11, a12, a13, b1 := sumX4, sumX3, sumX2, sumX2Y
	a21, a22, a23, b2 := sumX3, sumX2, sumX, sumXY
	a31, a32, a33, b3 := sumX2, sumX, n, sumY

	d := det(
		a11, a12, a13,
		a21, a22, a23,
		a31, a32, a33)
	detA := det(
		b1, a12, a13,
		b2, a22, a23,
		b3, a32, a33)
	detB := det(
		a11, b1, a13,
		a21, b2, a23,
		a31, b3, a33)
	detC := det(
		a11, a12, b1,
		a21, a22, b2,
		a31, a32, b3)

	a = detA / d
	b = detB / d
	c = detC / d
	return
}

func main() {
	pts := []Point{{400, 200}, {500, 300}, {300, 300}, {400, 500}}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := NewGame(screenWidth, screenHeight, pts)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
