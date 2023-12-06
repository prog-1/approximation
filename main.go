package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"

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

type Game struct {
	width, height int
	p             []Point
}

func (g *Game) Update() error {
	return nil
}

func approximation(p []Point) (a, b float64) {
	var sXY, sX, sY, ssqX float64
	n := float64(len(p))
	for _, i := range p {
		sXY += i.x * i.y
		sX += i.x
		sY += i.y
		ssqX += i.x * i.x
	}
	a = (n*sXY - sX*sY) / (n*ssqX - sX*sX)
	b = (sY - a*sX) / n
	return a, b
}

func (g *Game) Draw(screen *ebiten.Image) {
	a, b := approximation(g.p)
	vector.StrokeLine(screen, 0, float32(b), float32(g.width), float32(a*float64(g.width)+b), 2, color.RGBA{250, 250, 0, 0}, false)
	for _, i := range g.p {
		vector.DrawFilledCircle(screen, float32(i.x), float32(i.y), 2, color.RGBA{255, 0, 255, 255}, false)
	}
}

func (g *Game) Layout(outWidth, outHeight int) (w, h int) {
	return g.width, g.height
}

func NewGame(width, height int, p []Point) *Game {
	return &Game{
		width:  width,
		height: height,
		p:      p,
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := NewGame(screenWidth, screenHeight, []Point{{20, 430}, {50, 20}, {320, 40}})
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
